package log

func Init() {
	DefaultLogger = NewZapLogger()
	defaultHelper = NewHelper(DefaultLogger)
}

func Defer() {
	DefaultLogger.Defer()
}
