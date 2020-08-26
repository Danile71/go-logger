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

func Error(args ...interface{}) {
	Log(ERR, args...)
}

func Errorf(format string, args ...interface{}) {
	Logf(ERR, format, args...)
}
