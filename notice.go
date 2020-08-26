package logger

func Notice(args ...interface{}) {
	Log(NOTICE, args)
}

func Noticef(format string, args ...interface{}) {
	Logf(NOTICE, format, args)
}
