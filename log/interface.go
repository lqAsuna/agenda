package log4go

type Log4go struct {
}

func (this *Log4go) Println(enable bool, lever int, mode int, message ...interface{}) {
	Println(enable, lever, mode, message...)
}

func (this *Log4go) Printf(enable bool, lever int, mode int, format string, message ...interface{}) {
	Printf(enable, lever, mode, format, message...)
}
