package logger

import (
	"encoding/json"
	"fmt"
	"os"
	"runtime"
	"time"
)

type Level int

var debugLevel = EMERG
var api string

const (
	EMERG Level = iota
	ALERT
	CRIT
	ERR
	WARNING
	NOTICE
	INFO
	DEBUG
)

type LogMessage struct {
	Timestamp string `json:"time"`
	Level     Level  `json:"level"`
	Message   string `json:"message"`
}

func init() {
	level := os.Getenv("DEBUG_LEVEL")
	switch level {
	case "0":
		debugLevel = EMERG
	case "1":
		debugLevel = ALERT
	case "2":
		debugLevel = CRIT
	case "3":
		debugLevel = ERR
	case "4":
		debugLevel = WARNING
	case "5":
		debugLevel = NOTICE
	case "6":
		debugLevel = INFO
	case "7":
		debugLevel = DEBUG
	default:
		debugLevel = ERR
	}

	api = os.Getenv("DEBUG_URL")

}

func Log(level Level, args ...interface{}) LogMessage {
	_, fn, line, _ := runtime.Caller(1)
	message := fmt.Sprint(args...)
	msg := fmt.Sprintf("%s:%d|%s", fn, line, message)
	m := LogMessage{Timestamp: time.Now().String(), Level: level, Message: msg}
	b, err := json.Marshal(m)
	if err == nil {
		if level <= debugLevel {
			fmt.Println(string(b))
			go send(b)
		}
	}
	return m
}

func Logf(level Level, format string, args ...interface{}) LogMessage {
	_, fn, line, _ := runtime.Caller(1)
	message := fmt.Sprintf(format, args...)
	msg := fmt.Sprintf("%s:%d|%s", fn, line, message)
	m := LogMessage{Timestamp: time.Now().String(), Level: level, Message: msg}
	b, err := json.Marshal(m)
	if err == nil {
		if level <= debugLevel {
			fmt.Println(string(b))
			go send(b)
		}
	}
	return m
}
