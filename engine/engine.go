package engine

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

func (e *Engine) Out() io.Writer {
	return e.out
}

func (e *Engine) Level() LogLevel {
	return e.level
}

func (e *Engine) Encoder() Encoder {
	return e.enc
}

func DefaultEngine() *Engine {
	return &Engine{
		level: FATAL,
		enc:   DefaultTextEncoder(),
		out:   os.Stdout,
	}
}
