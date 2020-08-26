package logger

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

func OnError(err error) bool {
	if err == nil {
		return false
	}

	_, fn, line, _ := runtime.Caller(1)
	msg := fmt.Sprintf("%s:%d|%v", fn, line, err)
	m := LogMessage{Timestamp: time.Now().String(), Level: ERR, Message: msg}
	b, err := json.Marshal(m)
	if err == nil {
		if 3 <= debugLevel {
			fmt.Println(string(b))
		}
	}
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
