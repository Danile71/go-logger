package logger

func Warn(args ...interface{}) LogMessage {
	return Log(WARNING, args...)
}

func Warnf(format string, args ...interface{}) LogMessage {
	return Logf(WARNING, format, args...)
}
