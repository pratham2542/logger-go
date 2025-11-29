package logger

import (
	"io"
	"os"
)

type Engine struct {
	level LogLevel
	enc   Encoder
	out   io.Writer
}

func NewEngine(level LogLevel, encoder Encoder, output io.Writer) *Engine {
	return &Engine{
		level: level,
		enc:   encoder,
		out:   output,
	}
}

func DefaultEngine() *Engine {
	return &Engine{
		level: FATAL,
		enc:   DefaultTextEncoder(),
		out:   os.Stdout,
	}
}
