package logger

import (
	"runtime"
)

func Alert(args ...interface{}) LogMessage {
	_, fn, line, _ := runtime.Caller(1)
	return log(ALERT, fn, line, args...)
}

func Alertf(format string, args ...interface{}) LogMessage {
	_, fn, line, _ := runtime.Caller(1)
	return logf(ALERT, fn, line, format, args...)
}
