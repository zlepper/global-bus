package global_bus

import (
	"fmt"
	"log"
)

type LogLevel byte

const (
	Trace LogLevel = iota
	Debug
	Information
	Error
	Critical
)

func (l LogLevel) String() string {
	switch l {
	case Trace:
		return "trace"
	case Debug:
		return "debug"
	case Information:
		return "information"
	case Error:
		return "error"
	case Critical:
		return "critical"
	default:
		return "UNKNOWN"
	}
}

type Logger interface {
	Log(level LogLevel, format string, v ...interface{})
	Panic(format string, v ...interface{})
}

type StdLogger struct {
}

func (s StdLogger) Log(level LogLevel, format string, v ...interface{}) {
	log.Printf("[%s] %s", level.String(), fmt.Sprintf(format, v...))
}

func (s StdLogger) Panic(format string, v ...interface{}) {
	log.Panicf(format, v...)
}

type betterLogger struct {
	Logger
}

func (l betterLogger) Trace(format string, v ...interface{}) {
	l.Log(Trace, format, v...)
}
func (l betterLogger) Debug(format string, v ...interface{}) {
	l.Log(Debug, format, v...)
}
func (l betterLogger) Info(format string, v ...interface{}) {
	l.Log(Information, format, v...)
}
func (l betterLogger) Error(format string, v ...interface{}) {
	l.Log(Error, format, v...)
}
func (l betterLogger) Critical(format string, v ...interface{}) {
	l.Log(Critical, format, v...)
}
