package main

import "go.uber.org/zap"

var sugarLogger *zap.SugaredLogger


func main() {
	//newExample()
	//newDevelopment()
	//newProduction()

	//record()
	//sugared()

	// 自定义logger
	//initLogger()
	initLogger1()
	sugarLogger.Info("this is info message")
	sugarLogger.Infof("this is %s, %d", "aaa", 1234)
	sugarLogger.Error("this is error message")
	sugarLogger.Info("this is info message")
}
