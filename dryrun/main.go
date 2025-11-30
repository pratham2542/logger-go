package main

import (
	"os"
	"time"

	"github.com/pratham2542/logger-go"
	eng "github.com/pratham2542/logger-go/engine"
)

func main() {

	engine := eng.NewEngine(eng.DEBUG, eng.DefaultTextEncoder(), os.Stdout)
	l := logger.NewLogger(engine)
	i := 0
	for {
		l.Info("logger loop started")
		l.Debug("", eng.Int("Debug value", i))
		i++
		time.Sleep(1000 * time.Millisecond)
	}
}
