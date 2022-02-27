package main

import "go.uber.org/zap"

func test() {
	zap.NewProductionConfig()
}
