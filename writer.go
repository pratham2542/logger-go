package logger

import (
	"os"
	"runtime"

	fastBuffer "github.com/pratham2542/logger-go/buffer"
	"github.com/pratham2542/logger-go/engine"
)

// Public API
func (l *Logger) Debug(msg string, fields ...engine.Field) { l.log(engine.DEBUG, msg, fields...) }
func (l *Logger) Info(msg string, fields ...engine.Field)  { l.log(engine.INFO, msg, fields...) }
func (l *Logger) Warn(msg string, fields ...engine.Field)  { l.log(engine.WARN, msg, fields...) }
func (l *Logger) Error(msg string, fields ...engine.Field) { l.log(engine.ERROR, msg, fields...) }
func (l *Logger) Fatal(msg string, fields ...engine.Field) {
	l.log(engine.FATAL, msg, fields...)
	os.Exit(1)
}

func (l *Logger) writeEntry(e *engine.Entry) {
	if e.Level < l.engine.Level() {
		return
	}

	buf := l.bufPool.Get().(*fastBuffer.FastBuffer)
	buf.Reset()

	buf.Append(l.tsCache.Bytes())
	buf.AppendByte(' ')

	l.Encoder().Encode(e, buf)

	_, _ = l.out.Write(buf.Bytes())

	l.bufPool.Put(buf)
}

func (l *Logger) log(level engine.LogLevel, msg string, fields ...engine.Field) {
	file, line := l.captureCallerIfEnabled()
	e := engine.NewEntry(level, msg, fields, file, line)
	l.writeEntry(e)
}

func (l *Logger) captureCallerIfEnabled() (string, int) {
	if !l.withCaller {
		return "", 0
	}
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}
	return file, line
}
