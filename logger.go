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
	Level
	Format
	TimeFormat string
}

var (
	def     = Output{os.Stderr, EMERG, JSON, ""}
	outputs = []Output{def}
)

const (
	fformat = "%s:%d|%s"
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

func AddOutput(out io.Writer, level Level, format ...Format) Output {
	output := Output{out, level, JSON, ""}
	if len(format) > 0 {
		output.Format = format[0]
	}
	outputs = append(outputs, output)
	return output
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
	def.Level = level
}

func (output *Output) SetLevel(level Level) {
	output.Level = level
}

func SetTimeFormat(format string) {
	def.TimeFormat = format
}

func (output *Output) SetTimeFormat(format string) {
	output.TimeFormat = format
}

func init() {
	level := os.Getenv("DEBUG_LEVEL")
	switch level {
	case "0":
		def.Level = EMERG
	case "1":
		def.Level = ALERT
	case "2":
		def.Level = CRIT
	case "3":
		def.Level = ERR
	case "4":
		def.Level = WARNING
	case "5":
		def.Level = NOTICE
	case "6":
		def.Level = INFO
	case "7":
		def.Level = DEBUG
	default:
		def.Level = ERR
	}

	def.TimeFormat = os.Getenv("DEBUG_TIME_FORMAT")
}

func llog(level Level, msg string) LogMessage {
	m := LogMessage{Timestamp: time.Now().String(), Level: level, Message: msg}

	for _, out := range outputs {
		if level <= out.Level {
			if out.TimeFormat != "" {
				m.Timestamp = time.Now().Format(out.TimeFormat)
			}

			switch out.Format {
			case STRING:
				out.Write([]byte(m.String() + "\n"))
			default:
				// yes, we skipped error
				json.NewEncoder(out).Encode(m)
			}

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
