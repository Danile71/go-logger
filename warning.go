package logger

func Warn(args ...interface{}) {
	Log(WARNING, args...)
}

func Warnf(format string, args ...interface{}) {
	Logf(WARNING, format, args...)
}
