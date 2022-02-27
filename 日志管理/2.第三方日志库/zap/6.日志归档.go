package main

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)


// 日志分割
func getLogWriter1() zapcore.WriteSyncer {
	hook := &lumberjack.Logger{
		Filename:   "./test1.log",
		MaxSize:    10,
		MaxBackups: 5,
		MaxAge:     30,
		Compress:   false,
	}
	return zapcore.AddSync(hook)
}


func initLogger1() {
	encoder := getEncoder()
	writeSyncer := getLogWriter1()

	// 需要三个配置，编码器，写入位置，日志级别
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	// 新建自定义的logger对象
	// zap.AddCaller()  添加将调用函数信息记录到日志中的功能
	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}
