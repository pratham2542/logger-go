package main

import (
	"os"
	"time"

	"github.com/pratham2542/logger-go"
)

func main() {

	engine := logger.NewEngine(logger.DEBUG, logger.DefaultTextEncoder(), os.Stdout)
	l := logger.NewLogger(engine)
	i := 0
	for {
		l.Info("logger loop started")
		l.Debug("", logger.Int("Debug value", i))
		i++
		time.Sleep(1000 * time.Millisecond)
	}
}
