package logger

import (
	"runtime"
)

func Emerg(args ...interface{}) LogMessage {
	_, fn, line, _ := runtime.Caller(1)
	return log(EMERG, fn, line, args...)
}

func Emergf(format string, args ...interface{}) LogMessage {
	_, fn, line, _ := runtime.Caller(1)
	return logf(EMERG, fn, line, format, args...)
}
