package logger

import (
	"runtime"
)

func Warn(args ...interface{}) LogMessage {
	_, fn, line, _ := runtime.Caller(1)
	return log(WARNING, fn, line, args...)
}

func Warnf(format string, args ...interface{}) LogMessage {
	_, fn, line, _ := runtime.Caller(1)
	return logf(WARNING, fn, line, format, args...)
}
