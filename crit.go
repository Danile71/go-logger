package logger

import (
	"runtime"
)

func Crit(args ...interface{}) LogMessage {
	_, fn, line, _ := runtime.Caller(1)
	return log(CRIT, fn, line, args...)
}

func Critf(format string, args ...interface{}) LogMessage {
	_, fn, line, _ := runtime.Caller(1)
	return logf(CRIT, fn, line, format, args...)
}
