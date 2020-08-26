package logger

func Debug(args ...interface{}) LogMessage {
	return Log(DEBUG, args...)
}

func Debugf(format string, args ...interface{}) LogMessage {
	return Logf(DEBUG, format, args...)
}
