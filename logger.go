package logger

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"runtime"
	"time"
)

// Log level
type Level int

// Output format
type Format int

// Output instance
type Output struct {
	io.Writer
	Format
}

var (
	debugLevel = EMERG
	timeFormat string
	fformat    = "%s:%d|%s"
	outputs    = []Output{Output{os.Stderr, JSON}}
)

const (
	JSON Format = iota
	STRING
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

func AddOutput(out io.Writer, format ...Format) {
	output := Output{out, JSON}
	if len(format) > 0 {
		output.Format = format[0]
	}
	outputs = append(outputs, output)
}

type LogMessage struct {
	Timestamp string `json:"time"`
	Level     Level  `json:"level"`
	Message   string `json:"message"`
}

func (message *LogMessage) String() string {
	return fmt.Sprintf("%s %s", message.Timestamp, message.Message)
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
		for _, out := range outputs {
			//if formatOutput ==
			// yes, we skipped error
			json.NewEncoder(out).Encode(m)
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
