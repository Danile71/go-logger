package logger

func Warning(args ...interface{}) {
	Log(WARNING, args)
}

func Warningf(format string, args ...interface{}) {
	Logf(WARNING, format, args)
}
