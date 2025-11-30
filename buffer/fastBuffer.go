package fastBuffer

import (
	"strconv"
	"sync"
)

type FastBuffer struct {
	b []byte
}

func NewFastBuffer(b []byte) *FastBuffer {
	return &FastBuffer{b}
}

func (fb *FastBuffer) Reset() {
	fb.b = fb.b[:0]
}

func (fb *FastBuffer) Bytes() []byte {
	return fb.b
}

func (fb *FastBuffer) Append(p []byte) (int, error) {
	fb.b = append(fb.b, p...)
	return len(p), nil
}

func (fb *FastBuffer) AppendString(s string) (int, error) {
	fb.b = append(fb.b, s...)
	return len(s), nil
}

func (fb *FastBuffer) AppendInt(i int) (int, error) {
	old := len(fb.b)
	fb.b = strconv.AppendInt(fb.b, int64(i), 10)
	return len(fb.b) - old, nil
}

func (fb *FastBuffer) AppendFloat(f float64) (int, error) {
	old := len(fb.b)
	fb.b = strconv.AppendFloat(fb.b, f, 'g', -1, 64)
	return len(fb.b) - old, nil
}

func (fb *FastBuffer) AppendBool(v bool) (int, error) {
	old := len(fb.b)
	fb.b = strconv.AppendBool(fb.b, v)
	return len(fb.b) - old, nil
}

func (fb *FastBuffer) AppendByte(c byte) error {
	fb.b = append(fb.b, c)
	return nil
}

// trim before put if the buffer grows beyond a limit (64 KB)
func (fb *FastBuffer) TrimAndPut(pool *sync.Pool) {
	if cap(fb.b) > 64*1024 { // >64KB, too big
		pool.Put(&FastBuffer{b: make([]byte, 0, 1024)}) // replaced current buffer with fresh small buffer
	} else {
		fb.Reset()
		pool.Put(fb)
	}
}
