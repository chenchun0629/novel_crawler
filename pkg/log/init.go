package log

func Init() {
	DefaultLogger = NewZapLogger()
}
