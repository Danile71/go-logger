package logger

func Alert(args ...interface{}) LogMessage {
	return Log(ALERT, args...)
}

func Alertf(format string, args ...interface{}) LogMessage {
	return Logf(ALERT, format, args...)
}
