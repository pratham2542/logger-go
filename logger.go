package logger

import (
	"bytes"
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
		out:      &lockedWriter{w: out}, // lockeedWrite will make sure that no writes are getting intertwined
		bufPool: sync.Pool{
			New: func() interface{} {
				// This is the only buffer creation
				// After this initialization logger will reuse it again and again
				return new(bytes.Buffer) // create a new buffer in pool
			},
		},
	}
}
