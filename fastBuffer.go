package logger

import "sync"

type fastBuffer struct {
	b []byte
}

func (fb *fastBuffer) Reset() {
	fb.b = fb.b[:0]
}

func (fb *fastBuffer) Bytes() []byte {
	return fb.b
}

func (fb *fastBuffer) Write(p []byte) (int, error) {
	fb.b = append(fb.b, p...)
	return len(p), nil
}

func (fb *fastBuffer) WriteString(s string) (int, error) {
	fb.b = append(fb.b, s...)
	return len(s), nil
}

func (fb *fastBuffer) WriteByte(c byte) error {
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
