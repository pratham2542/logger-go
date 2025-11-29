package logger

import (
	"io"
	"os"
	"sync"
)

type Logger struct {
	minLevel   LogLevel
	fullPath   bool
	withCaller bool
	out        io.Writer
	bufPool    sync.Pool
	tsCache    *timestampCache
	fileCache  sync.Map // cache for base filenames
	encoder    Encoder
}

// Synchronous logger
func NewLogger(level LogLevel, out io.Writer, fullPath, withCaller bool) *Logger {
	if out == nil {
		out = os.Stdout
	}
	return &Logger{
		minLevel:   level,
		fullPath:   fullPath,
		withCaller: withCaller,
		out:        &lockedWriter{w: out}, // lockedWriter will make sure that no writes are getting intertwined
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
		tsCache:   newTimestampCache(),
		fileCache: sync.Map{},
		encoder:   defaultTextEncoder(),
	}
}
