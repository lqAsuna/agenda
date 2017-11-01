package log4go

import (
	"log"
	"runtime/debug"

	"xxx/loggor"
)

const (
	_TABLE_ = "\t" 
)

const (
	LEVER_UNKNOW = 0
	LEVER_DEBUG  = 1
	LEVER_INFO   = 2
	LEVER_WARN   = 3
	LEVER_ERROR  = 4
	LEVER_FATAL  = 5
)
const (
	MODE_STDIO = 1
	MODE_FILE  = 2
	MODE_NET   = 3
)
const (
	LOG_TYPE_UNKNOW = 1000
	LOG_TYPE_DEBUG  = 1001
	LOG_TYPE_INFO   = 1002
	LOG_TYPE_WARN   = 1003
	LOG_TYPE_ERROR  = 1004
	LOG_TYPE_FATAL  = 1005
)
const (
	LOG_TYPE_ROOT_NAME   = "./logs/level/"
	LOG_TYPE_UNKNOW_NAME = "unknow"
	LOG_TYPE_DEBUG_NAME  = "debug"
	LOG_TYPE_INFO_NAME   = "info"
	LOG_TYPE_WARN_NAME   = "warn"
	LOG_TYPE_ERROR_NAME  = "error"
	LOG_TYPE_FATAL_NAME  = "fatal"
)

var debugLoger *loggor.Logger
var infoLoger *loggor.Logger 
var warnLoger *loggor.Logger
var errorLoger *loggor.Logger
var fatalLoger *loggor.Logger 

var (
	IS_DEBUG = false
)

func SetDebug(debug bool) {
	IS_DEBUG = debug
}

func _init_() {
	debugLoger = &loggor.Logger{}
	infoLoger = &loggor.Logger{}
	warnLoger = &loggor.Logger{}
	errorLoger = &loggor.Logger{}
	fatalLoger = &loggor.Logger{}

	debugLoger.SetDebug(IS_DEBUG)
	debugLoger.SetType(LOG_TYPE_DEBUG)
	debugLoger.SetRollingFile(LOG_TYPE_ROOT_NAME, LOG_TYPE_DEBUG_NAME, 5, 100, loggor.MB)

	warnLoger.SetDebug(IS_DEBUG)
	infoLoger.SetType(LOG_TYPE_INFO)
	infoLoger.SetRollingFile(LOG_TYPE_ROOT_NAME, LOG_TYPE_INFO_NAME, 5, 100, loggor.MB)

	warnLoger.SetDebug(IS_DEBUG)
	warnLoger.SetType(LOG_TYPE_WARN)
	warnLoger.SetRollingFile(LOG_TYPE_ROOT_NAME, LOG_TYPE_WARN_NAME, 5, 100, loggor.MB)

	errorLoger.SetDebug(IS_DEBUG)
	errorLoger.SetType(LOG_TYPE_ERROR)
	errorLoger.SetRollingFile(LOG_TYPE_ROOT_NAME, LOG_TYPE_ERROR_NAME, 5, 100, loggor.MB)

	fatalLoger.SetDebug(IS_DEBUG)
	fatalLoger.SetType(LOG_TYPE_FATAL)
	fatalLoger.SetRollingFile(LOG_TYPE_ROOT_NAME, LOG_TYPE_FATAL_NAME, 5, 100, loggor.MB)
}


func LogServer() {
	_init_()
}

func Println(enable bool, lever int, mode int, message ...interface{}) {
	if !enable {
		return
	}

	defer func() {
		if e, ok := recover().(error); ok {
			log.Println("ERR: panic in %s - %v", message, e)
			log.Println(string(debug.Stack()))
		}
	}()

	switch mode {
	case MODE_STDIO:
		log.Println(message)
		break
	case MODE_NET:
		break
	case MODE_FILE:
		if nil == infoLoger {
			_init_()
		}
		switch lever {
		case LEVER_DEBUG:
			(*debugLoger).Println(message)
			break
		case LEVER_INFO:
			(*infoLoger).Println(message)
			break
		case LEVER_WARN:
			(*warnLoger).Println(message)
			break
		case LEVER_ERROR:
			(*errorLoger).Println(message)
			break
		case LEVER_FATAL:
			(*fatalLoger).Println(message)
			break
		default:
			break
		}
		break
	default:
		break
	}
}

func Printf(enable bool, lever int, mode int, format string, message ...interface{}) {
	if !enable {
		return
	}

	defer func() {
		if e, ok := recover().(error); ok {
			log.Println("ERR: panic in %s - %v", message, e)
			log.Println(string(debug.Stack()))
		}
	}()
	switch mode {
	case MODE_STDIO:
		log.Println(format, message)
		break
	case MODE_NET:
		//网络方式
		break
	case MODE_FILE:
		if nil == infoLoger {
			_init_()
		}
		switch lever {
		case LEVER_DEBUG:
			(*debugLoger).Println(format, message)
			break
		case LEVER_INFO:
			(*infoLoger).Println(format, message)
			break
		case LEVER_WARN:
			(*warnLoger).Println(format, message)
			break
		case LEVER_ERROR:
			(*errorLoger).Println(format, message)
			break
		case LEVER_FATAL:
			(*fatalLoger).Println(format, message)
			break
		default:
			break
		}
		break
	default:
		break
	}
}
