package engine

import fastBuffer "github.com/pratham2542/logger-go/buffer"

// Encoder formats a log Entry into a fastBuffer.

type Encoder interface {
	Encode(e *Entry, buf *fastBuffer.FastBuffer)
}

type textEncoder struct{}

func DefaultTextEncoder() Encoder { return &textEncoder{} }

func (t *textEncoder) Encode(e *Entry, buf *fastBuffer.FastBuffer) {
	buf.AppendString(e.Level.String())
	buf.AppendByte(' ')
	buf.AppendString(e.Msg)

	for _, f := range e.Fields {
		buf.AppendByte(' ')
		buf.AppendString(f.Key)
		buf.AppendByte('=')
		f.AppendValueTo(buf)
	}

	buf.AppendByte('\n')
}
