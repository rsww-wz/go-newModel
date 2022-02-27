package main

import (
	"go.uber.org/zap"
)

// 记录日志
func record() {
	log, _ := zap.NewDevelopment()

	log.Info("初级日志", zap.Int("Python", 1))
	log.Info("初级日志", zap.Int("C", 2))
	log.Info("初级日志", zap.Int("Java", 3))
	log.Info("初级日志", zap.Int("Go", 19))
	log.Info("\n")

	log.Info("警告日志",
		zap.Any("name","rs"),
		zap.Any("age",21),
		zap.Any("gender",false))
}
