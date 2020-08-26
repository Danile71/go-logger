package logger

func Info(args ...interface{}) LogMessage {
	return Log(INFO, args...)
}

func Infof(format string, args ...interface{}) LogMessage {
	return Logf(INFO, format, args...)
}
