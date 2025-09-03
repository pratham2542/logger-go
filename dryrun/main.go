package main

import (
	"os"
	"time"

	"github.com/pratham2542/logger-go"
)

func main() {
	logger := logger.NewLogger(logger.DEBUG, os.Stdout, true, true)
	i := 0
	for {
		logger.Info("logger loop started")
		logger.Debug("Debug value:", i)
		i++
		time.Sleep(1000 * time.Millisecond)
	}
}
