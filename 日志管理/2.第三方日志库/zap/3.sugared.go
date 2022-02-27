package main

import "go.uber.org/zap"

func sugared() {
	logger,_:=zap.NewDevelopment()
	slogger:=logger.Sugar()

	slogger.Debugf("debug message age is %d,agender is %s",19,"man")
	slogger.Info("Info() uses sprint")
	slogger.Infof("Infof() uses %s","sprintf")
	slogger.Infow("Infow() allows tags","name","Legolas","type",1)

	logger = slogger.Desugar()
	logger.Debug("打印变量",zap.Int("循环变量",1))
}