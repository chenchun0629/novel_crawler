package xgo

func Init() {

}

func Close() {
	defaultManager.closeAll()
}
