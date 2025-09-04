package main

import (
	"os"
	"time"

	"github.com/pratham2542/logger-go"
)

func main() {
	l := logger.NewLogger(logger.DEBUG, os.Stdout, false, false)
	i := 0
	for {
		l.Info("logger loop started")
		l.Debug("", logger.Int("Debug value", i))
		i++
		time.Sleep(1000 * time.Millisecond)
	}
}
