package global_bus

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type logEvent struct {
	level   LogLevel
	message string
}

type innerTestLogger struct {
	messages []logEvent
	panics   []string
}

func (l *innerTestLogger) Log(level LogLevel, format string, v ...interface{}) {
	l.messages = append(l.messages, logEvent{
		level:   level,
		message: fmt.Sprintf(format, v...),
	})
}

func (l *innerTestLogger) Panic(format string, v ...interface{}) {
	l.panics = append(l.panics, fmt.Sprintf(format, v...))
	panic(fmt.Sprintf(format, v...))
}

func createTestLogger() (*innerTestLogger, betterLogger) {
	inner := &innerTestLogger{
		messages: []logEvent{},
		panics:   []string{},
	}

	return inner, betterLogger{inner}
}

func TestBetterLogger_Critical(t *testing.T) {
	inner, logger := createTestLogger()

	logger.Critical("foo: %d", 42)

	assert.Len(t, inner.messages, 1)
	assert.Equal(t, inner.messages[0], logEvent{level: Critical, message: "foo: 42"})
}

func TestBetterLogger_Error(t *testing.T) {
	inner, logger := createTestLogger()

	logger.Error("foo: %d", 42)

	assert.Len(t, inner.messages, 1)
	assert.Equal(t, inner.messages[0], logEvent{level: Error, message: "foo: 42"})
}
func TestBetterLogger_Information(t *testing.T) {
	inner, logger := createTestLogger()

	logger.Info("foo: %d", 42)

	assert.Len(t, inner.messages, 1)
	assert.Equal(t, inner.messages[0], logEvent{level: Information, message: "foo: 42"})
}
func TestBetterLogger_Debug(t *testing.T) {
	inner, logger := createTestLogger()

	logger.Debug("foo: %d", 42)

	assert.Len(t, inner.messages, 1)
	assert.Equal(t, inner.messages[0], logEvent{level: Debug, message: "foo: 42"})
}
func TestBetterLogger_Trace(t *testing.T) {
	inner, logger := createTestLogger()

	logger.Trace("foo: %d", 42)

	assert.Len(t, inner.messages, 1)
	assert.Equal(t, inner.messages[0], logEvent{level: Trace, message: "foo: 42"})
}

func TestBetterLogger_Panic(t *testing.T) {
	inner, logger := createTestLogger()

	assert.PanicsWithValue(t, "foo: 42", func() {
		logger.Panic("foo: %d", 42)

	})

	assert.Len(t, inner.panics, 1)
	assert.Equal(t, inner.panics[0], "foo: 42")
}

func TestLogLevel_String(t *testing.T) {
	cases := []struct {
		level LogLevel
		text  string
	}{
		{
			Critical,
			"critical",
		},
		{
			Error,
			"error",
		},
		{
			Information,
			"information",
		},
		{
			Debug,
			"debug",
		},
		{
			Trace,
			"trace",
		},
		{
			42,
			"UNKNOWN",
		},
	}

	for _, s := range cases {
		assert.Equal(t, s.text, s.level.String())
	}
}
