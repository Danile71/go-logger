package logger

import (
	"runtime"
)

func Notice(args ...interface{}) LogMessage {
	_, fn, line, _ := runtime.Caller(1)
	return log(NOTICE, fn, line, args...)
}

func Noticef(format string, args ...interface{}) LogMessage {
	_, fn, line, _ := runtime.Caller(1)
	return logf(NOTICE, fn, line, format, args...)
}
