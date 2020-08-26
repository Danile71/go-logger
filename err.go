package logger

import (
	"runtime"
)

func OnError(err error) bool {
	if err == nil {
		return false
	}
	_, fn, line, _ := runtime.Caller(1)
	log(ERR, fn, line, err)
	return true
}

func Error(args ...interface{}) LogMessage {
	_, fn, line, _ := runtime.Caller(1)
	return log(ERR, fn, line, args...)
}

func Errorf(format string, args ...interface{}) LogMessage {
	_, fn, line, _ := runtime.Caller(1)
	return logf(ERR, fn, line, format, args...)
}
