package logger

func Crit(args ...interface{}) LogMessage {
	return Log(CRIT, args...)
}

func Critf(format string, args ...interface{}) LogMessage {
	return Logf(CRIT, format, args...)
}
