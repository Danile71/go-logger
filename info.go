package logger

import (
	"runtime"
)

func Info(args ...interface{}) LogMessage {
	_, fn, line, _ := runtime.Caller(1)
	return log(INFO, fn, line, args...)
}

func Infof(format string, args ...interface{}) LogMessage {
	_, fn, line, _ := runtime.Caller(1)
	return logf(INFO, fn, line, format, args...)
}
