package logger

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

type Level int

var (
	debugLevel = EMERG
	timeFormat string
	fformat    = "%s:%d|%s"
	outputs    = []io.Writer{os.Stderr}
)

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

func AddOutput(out io.Writer) {
	outputs = append(outputs, out)
}

type LogMessage struct {
	Timestamp string `json:"time"`
	Level     Level  `json:"level"`
	Message   string `json:"message"`
}

func SetLevel(level Level) {
	debugLevel = level
}

func SetTimeFormat(format string) {
	timeFormat = format
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

	timeFormat = os.Getenv("DEBUG_TIME_FORMAT")
}

func llog(level Level, msg string) LogMessage {
	m := LogMessage{Timestamp: time.Now().String(), Level: level, Message: msg}
	if timeFormat != "" {
		m.Timestamp = time.Now().Format(timeFormat)
	}
	if level <= debugLevel {
		// copy log message to all outputs through bytes.Buffer
		buf := &bytes.Buffer{}
		_ = json.NewEncoder(buf).Encode(m)
		for _, out := range outputs {
			go io.Copy(out, buf)
		}
	}
	return m
}

func log(level Level, fn string, line int, args ...interface{}) LogMessage {
	return llog(level, fmt.Sprintf(fformat, fn, line, fmt.Sprint(args...)))
}

func logf(level Level, fn string, line int, format string, args ...interface{}) LogMessage {
	return llog(level, fmt.Sprintf(fformat, fn, line, fmt.Sprintf(format, args...)))
}

func Log(level Level, args ...interface{}) LogMessage {
	_, fn, line, _ := runtime.Caller(1)
	return log(level, fn, line, args...)
}

func Logf(level Level, format string, args ...interface{}) LogMessage {
	_, fn, line, _ := runtime.Caller(1)
	return logf(level, fn, line, format, args...)
}
