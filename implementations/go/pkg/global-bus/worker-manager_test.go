package global_bus

import (
	"context"
	"github.com/stretchr/testify/assert"
	"log"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type silentLogger struct {
}

func (s silentLogger) Log(level LogLevel, format string, v ...interface{}) {
	// Do nothing
}

func (s silentLogger) Panic(format string, v ...interface{}) {
	log.Panicf(format, v...)
}

// Tests that there will never be too many jobs processing at the same time
func TestWorkerProcessing(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	manager := newWorkerManager(silentLogger{}, ctx)

	defer cancel()

	workerCount := 4

	creatorCount := 20

	jobCount := 100

	manager.startWorkers(ctx, workerCount)

	var activeJobs int64

	jobCounts := make([]int64, 0)
	var lock sync.Mutex

	var wg sync.WaitGroup

	for i := 0; i < creatorCount; i++ {
		wg.Add(jobCount)
		go func() {

			for j := 0; j < jobCount; j++ {
				manager.process(func() {
					now := atomic.AddInt64(&activeJobs, 1)

					lock.Lock()
					defer lock.Unlock()
					jobCounts = append(jobCounts, now)

					atomic.AddInt64(&activeJobs, -1)
					wg.Done()
				})
			}
		}()
	}

	wg.Wait()

	var max int64

	for _, count := range jobCounts {
		if max < count {
			max = count
		}
	}

	assert.LessOrEqual(t, max, int64(workerCount))

}

func TestWorkerManager_ContextCancelShouldStopWorkers(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	manager := newWorkerManager(silentLogger{}, ctx)

	manager.startWorkers(ctx, 4)

	var beforeCount int64

	for i := 0; i < 10; i++ {
		manager.process(func() {
			atomic.AddInt64(&beforeCount, 1)
		})
	}

	assert.Equal(t, int64(10), beforeCount)

	cancel()

	var stopped bool

	go func() {
		manager.workersFinished.Wait()
		stopped = true
	}()

	time.Sleep(time.Microsecond)

	assert.True(t, stopped)
}
