package global_bus

import (
	"context"
	"sync"
)

type workerManager struct {
	workersFinished sync.WaitGroup
	workerChannel   chan func()
	logger          betterLogger
}

func newWorkerManager(logger Logger, ctx context.Context) *workerManager {
	c := make(chan func())

	go func() {
		<-ctx.Done()

		close(c)
	}()

	return &workerManager{
		workerChannel: c,
		logger:        betterLogger{logger},
	}
}

// Enqueue the callback on a worker process. The function returns
// when the callback has finished executing
func (m *workerManager) process(callback func()) {
	var wg sync.WaitGroup
	wg.Add(1)
	m.workerChannel <- func() {
		callback()
		wg.Done()
	}

	wg.Wait()
}

// Starts the specified number of workers
func (m *workerManager) startWorkers(ctx context.Context, count int) {
	m.logger.Debug("Starting %d Global Bus workers", count)
	for i := 0; i < count; i++ {
		m.startWorker(ctx)
	}
}

func (m *workerManager) startWorker(ctx context.Context) {
	m.logger.Debug("Starting Global Bus worker")
	m.workersFinished.Add(1)

	go func() {
		for {
			select {
			case <-ctx.Done():
				m.logger.Debug("Global Bus worker has detected a cancellation request. Stopping")
				m.workersFinished.Done()
				return
			case job := <-m.workerChannel:
				m.logger.Trace("New job arrived at Global Bus worker, executing...")
				job()
				m.logger.Trace("Worker finished job")
			}
		}
	}()

}
