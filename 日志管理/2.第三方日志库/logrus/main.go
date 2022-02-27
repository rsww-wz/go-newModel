package main

import (
	log "github.com/sirupsen/logrus"
	"os"
)

// 配置日志
func init() {
	// 设置日志格式为json格式
	log.SetFormatter(&log.JSONFormatter{})

	// 设置将日志输出到标准输出 (默认的输出为stderr,标准错误)
	//log.SetOutput(os.Stdout)

	// 日志消息输出可以是任意的io.Writer类型
	f, err := os.OpenFile("./logrus.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalln("文件创建/打开失败", err)
	}
	log.SetOutput(f)

	// 设置日志级别为warn以上
	log.SetLevel(log.WarnLevel)
}

func main() {
	//logrusTest()
	//logrusLevel()

	// 前面的函数已经退出程序了，所以不会执行后面的函数
	totalLogger()
}

// 一次性记录所有信息
func logrusLevel() {
	// 在TraceLevel级别以上的日志都会记录
	//log.SetLevel(log.TraceLevel)

	// 只会记录ErrorLevel级别及以上的日志
	log.SetLevel(log.ErrorLevel)

	log.Traceln("trace 日志")
	log.Debugln("debug 日志")
	log.Infoln("Info 日志")
	log.Warnln("warn 日志")
	log.Errorln("error msg")
	log.Fatalf("fatal 日志")
	log.Panicln("panic 日志")
}


// logrus推荐使用的方式记录日志
func logrusTest() {
	// 初始化中设置了warning级别及以上才会记录
	log.WithFields(log.Fields{
		"animal": "walrus",
		"size":   10,
	}).Info("a group of walrus emerges form the ocean")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 122,
	}).Warn("The group's number increased tremendously!")

	log.WithFields(log.Fields{
		"omg":    true,
		"number": 100,
	}).Fatalln("The ice breaks!")
}
