package logger

import (
	"io"
	"sync"

	fastBuffer "github.com/pratham2542/logger-go/buffer"
	eng "github.com/pratham2542/logger-go/engine"
)

type Logger struct {
	engine     *eng.Engine
	fullPath   bool
	withCaller bool
	callerSkip int
	out        io.Writer
	bufPool    sync.Pool
	tsCache    *timestampCache
	fileCache  sync.Map // cache for base filenames
}

// Synchronous logger
func NewLogger(engine *eng.Engine) *Logger {
	if engine == nil {
		engine = eng.DefaultEngine()
	}
	return &Logger{
		engine:     engine,
		withCaller: false,
		out:        NewLockedWriter(engine.Out()), // lockedWriter will make sure that no writes are getting intertwined
		bufPool: sync.Pool{
			New: func() any {
				// This is the only buffer creation
				// After this initialization logger will reuse it again and again
				// create a fixed sized byte array so no realloaction happens
				return fastBuffer.NewFastBuffer(make([]byte, 0, 1024)) // change it to ~4 KB if the log formate changes to json
			},
		},
		tsCache:    newTimestampCache(),
		fileCache:  sync.Map{},
		callerSkip: 0,
		fullPath:   false,
	}
}

func (l *Logger) Encoder() eng.Encoder {
	return l.engine.Encoder()
}
