package logger

func Alert(args ...interface{}) {
	Log(ALERT, args...)
}

func Alertf(format string, args ...interface{}) {
	Logf(ALERT, format, args...)
}
