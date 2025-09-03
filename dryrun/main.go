package main

import "github.com/pratham2542/logger-go"

func main() {
	// Console logger
	logger := logger.NewLogger(logger.DEBUG, nil)
	logger.Info("Starting application")
	logger.Debug("Debug value:", 42)
}
