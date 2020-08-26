package logger

func Emerg(args ...interface{}) LogMessage {
	return Log(EMERG, args...)
}

func Emergf(format string, args ...interface{}) LogMessage {
	return Logf(EMERG, format, args...)
}
