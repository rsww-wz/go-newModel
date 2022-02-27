package main

import "go.uber.org/zap"


func newExample() {
	log := zap.NewExample()

	log.Debug("this is debug message")
	log.Info("this is info message")
	log.Info("this is debug message",
		zap.Int("age", 24), zap.String("gender", "man"))

	log.Warn("this is warn message")
	log.Error("this is error message")
	log.Panic("this is panic message")
}


func newDevelopment() {
	log, _ := zap.NewDevelopment()
	log.Debug("this is debug message")
	log.Info("this is info message")
	log.Info("this is info message with fields",
		zap.Int("age", 24), zap.String("gender", "man"))

	log.Warn("this is warn message")
	log.Error("this is error message")
	log.DPanic("This is a DPANIC message")
	log.Panic("this is panic message")
	log.Fatal("This is a FATAL message")
}


func newProduction() {
	log, _ := zap.NewProduction()
	log.Debug("this is debug message")
	log.Info("this is info message")
	log.Info("this is info message with fields",
		zap.Int("age", 24), zap.String("gender", "man"))
	log.Warn("this is warn message")
	log.Error("this is error message")
	log.DPanic("This is a DPANIC message")
	log.Panic("this is panic message")
	log.Fatal("This is a FATAL message")
}
