package logutil

import (
	"github.com/jeanphorn/log4go"
)

// 目前的问题，不调用log4go.Close(),日志不会打印在console和文件中
func init() {
	log4go.LoadConfiguration("./log4go.xml", "xml")
}

func Info(arg0 interface{}, args ...interface{}) {
	if len(args) > 0 {
		log4go.Info(arg0, args)
	} else {
		log4go.Info(arg0)
	}
}
func Debug(arg0 interface{}, args ...interface{}) {
	if len(args) > 0 {
		log4go.Debug(arg0, args)
	} else {
		log4go.Debug(arg0)
	}
}
func Error(arg0 interface{}, args ...interface{}) {
	if len(args) > 0 {
		log4go.Error(arg0, args)
	} else {
		log4go.Error(arg0)
	}
}

func Warn(arg0 interface{}, args ...interface{}) {
	if len(args) > 0 {
		log4go.Warn(arg0, args)
	} else {
		log4go.Warn(arg0)
	}
}
