package logger

import (
	"strconv"
	"sync"
)

type fastBuffer struct {
	b []byte
}

func (fb *fastBuffer) Reset() {
	fb.b = fb.b[:0]
}

func (fb *fastBuffer) Bytes() []byte {
	return fb.b
}

func (fb *fastBuffer) Append(p []byte) (int, error) {
	fb.b = append(fb.b, p...)
	return len(p), nil
}

func (fb *fastBuffer) AppendString(s string) (int, error) {
	fb.b = append(fb.b, s...)
	return len(s), nil
}

func (fb *fastBuffer) AppendInt(i int) (int, error) {
	old := len(fb.b)
	fb.b = strconv.AppendInt(fb.b, int64(i), 10)
	return len(fb.b) - old, nil
}

func (fb *fastBuffer) AppendFloat(f float64) (int, error) {
	old := len(fb.b)
	fb.b = strconv.AppendFloat(fb.b, f, 'g', -1, 64)
	return len(fb.b) - old, nil
}

func (fb *fastBuffer) AppendBool(v bool) (int, error) {
	old := len(fb.b)
	fb.b = strconv.AppendBool(fb.b, v)
	return len(fb.b) - old, nil
}

func (fb *fastBuffer) AppendByte(c byte) error {
	fb.b = append(fb.b, c)
	return nil
}

// trim before put if the buffer grows beyond a limit (64 KB)
func (fb *fastBuffer) TrimAndPut(pool *sync.Pool) {
	if cap(fb.b) > 64*1024 { // >64KB, too big
		pool.Put(&fastBuffer{b: make([]byte, 0, 1024)}) // replaced current buffer with fresh small buffer
	} else {
		fb.Reset()
		pool.Put(fb)
	}
}
