package logger

func Notice(args ...interface{}) LogMessage {
	return Log(NOTICE, args...)
}

func Noticef(format string, args ...interface{}) LogMessage {
	return Logf(NOTICE, format, args...)
}
