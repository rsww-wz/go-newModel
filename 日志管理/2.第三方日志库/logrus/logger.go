package main

// logger是一种相对高级的用法, 对于一个大型项目, 往往需要一个全局的logrus实例，即logger对象来记录项目所有的日志

import (
	log "github.com/sirupsen/logrus"
	"os"
)

// logrus 提供了New函数来创建一个logrus实例
// 项目中，可以创建任意数量的logrus实例
var logger = log.New()

func totalLogger() {
	// 为当前logrus实例设置消息的输出，同样的
	// 可以设置logrus实例的输出到任意io.Writer
	logger.Out = os.Stderr

	// 为当前logrus实例设置消息输出格式为json格式
	// 也可以单独为某个logrus实例设置日志级别和hook
	logger.Formatter = &log.JSONFormatter{}

	// 日志级别设定到最低级别
	logger.Level = log.TraceLevel

	// 记录信息
	// 注意调用者，是全局的logger
	logger.WithFields(log.Fields{
		"animal":"walrus",
		"size":10,
	}).Info("A group of walrus emerges from the ocean")

	logger.WithFields(log.Fields{
		"stderr" : "控制台",
	}).Warn("你现在是在控制台输出日志")
}