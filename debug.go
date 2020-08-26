package logger

import (
	"runtime"
)

func Debug(args ...interface{}) LogMessage {
	_, fn, line, _ := runtime.Caller(1)
	return log(DEBUG, fn, line, args...)
}

func Debugf(format string, args ...interface{}) LogMessage {
	_, fn, line, _ := runtime.Caller(1)
	return logf(DEBUG, fn, line, format, args...)
}
