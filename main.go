package main

import (
	"github.com/hongjundu/logxi-example/logger"
)

func main() {

	logger.Init("example", "example.log", "/aaa", 10, 3, 30)

	logger.Debug("Debug Message", "name", "duhj")
	logger.Trace("Trace Message", "name", "duhj")
	logger.Info("Info Message", "name", "duhj")
	logger.Warn("Warn Message", "name", "duhj")
	logger.Error("Error Message", "name", "duhj")

	logger.Info("Exit")
}
