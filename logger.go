package logger

import (
	"io"
	"os"
	"sync"
)

type Logger struct {
	minLevel LogLevel
	out      io.Writer
	bufPool  sync.Pool
}

// Synchronous logger
func NewLogger(level LogLevel, out io.Writer) *Logger {
	if out == nil {
		out = os.Stdout
	}
	return &Logger{
		minLevel: level,
		out:      out,
		bufPool: sync.Pool{
			New: func() interface{} { return nil },
		},
	}
}
