package logger

func Info(args ...interface{}) {
	Log(INFO, args)
}

func Infof(format string, args ...interface{}) {
	Logf(INFO, format, args)
}
