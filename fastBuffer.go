package logger

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
