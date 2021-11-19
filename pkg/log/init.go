package log

func Init() {
	DefaultLogger = NewZapLogger()
}

func Defer() {
	DefaultLogger.Defer()
}
