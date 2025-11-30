package logger

import (
	"io"
	"sync"
)

// lockedWriter wraps an io.Writer with a mutex to prevent interleaving writes.
type lockedWriter struct {
	mu sync.Mutex
	w  io.Writer
}

func NewLockedWriter(w io.Writer) *lockedWriter {
	return &lockedWriter{
		w:  w,
		mu: sync.Mutex{},
	}
}

func (lw *lockedWriter) Write(p []byte) (int, error) {
	lw.mu.Lock()
	defer lw.mu.Unlock()
	return lw.w.Write(p)
}
