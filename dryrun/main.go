package main

import (
	"os"

	"github.com/pratham2542/logger-go"
)

func main() {
	// Console logger
	logger := logger.NewLogger(logger.DEBUG, os.Stdout, true)
	logger.Info("Starting application")
	logger.Debug("Debug value:", 42)
}
