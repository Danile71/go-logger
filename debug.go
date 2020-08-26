package logger

func Debug(args ...interface{}) {
	Log(DEBUG, args...)
}

func Debugf(format string, args ...interface{}) {
	Logf(DEBUG, format, args...)
}
