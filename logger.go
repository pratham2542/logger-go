package logger

import (
	"io"
	"sync"
)

type Logger struct {
	engine     *Engine
	fullPath   bool
	withCaller bool
	callerSkip int
	out        io.Writer
	bufPool    sync.Pool
	tsCache    *timestampCache
	fileCache  sync.Map // cache for base filenames
}

// Synchronous logger
func NewLogger(engine *Engine) *Logger {
	if engine == nil {
		engine = DefaultEngine()
	}
	return &Logger{
		engine:     engine,
		withCaller: false,
		out:        &lockedWriter{w: engine.out}, // lockedWriter will make sure that no writes are getting intertwined
		bufPool: sync.Pool{
			New: func() any {
				// This is the only buffer creation
				// After this initialization logger will reuse it again and again
				return &fastBuffer{ // create a new buffer in pool
					b: make([]byte, 0, 1024), // create a fixed sized byte array so no realloaction happens
					// change it to ~4 KB if the log formate changes to json
				}
			},
		},
		tsCache:    newTimestampCache(),
		fileCache:  sync.Map{},
		callerSkip: 0,
		fullPath:   false,
	}
}
