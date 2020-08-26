package logger

func Crit(args ...interface{}) {
	Log(CRIT, args)
}

func Critf(format string, args ...interface{}) {
	Logf(CRIT, format, args)
}
