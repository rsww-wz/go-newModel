package main

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)


// 编码器(如何写入日志)
func getEncoder() zapcore.Encoder {
	// 预先设置的编码器配置，是一个结构体
	encoderConfig := zap.NewProductionEncoderConfig()

	// 修改时间编码器，单独修改结构体里面的时间编码器
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// 在日志文件中使用大写字母记录日志级别
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// NewConsoleEncoder 打印更符合人们观察的方式
	return zapcore.NewConsoleEncoder(encoderConfig)
}


// 指定日志将写到那里去
func getLogWriter() zapcore.WriteSyncer {
	// 打开文件，获得文件指针
	file, _ := os.Create("./test.log")
	return zapcore.AddSync(file)
}


// 整合以上配置信息
func initLogger() {
	encoder := getEncoder()
	writeSyncer := getLogWriter()

	// 需要三个配置，编码器，写入位置，日志级别
	core := zapcore.NewCore(encoder, writeSyncer, zapcore.DebugLevel)

	// 新建自定义的logger对象
	// zap.AddCaller()  添加将调用函数信息记录到日志中的功能
	logger := zap.New(core, zap.AddCaller())
	sugarLogger = logger.Sugar()
}
