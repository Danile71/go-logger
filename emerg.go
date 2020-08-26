package logger

func Emerg(args ...interface{}) {
	Log(EMERG, args...)
}

func Emergf(format string, args ...interface{}) {
	Logf(EMERG, format, args...)
}
