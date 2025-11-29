package logger

import "time"

// Entry represents a single log record prepared for encoding.

type Entry struct {
	Time   time.Time
	Level  LogLevel
	Msg    string
	File   string
	Line   int
	Fields []Field
}

func newEntry(level LogLevel, msg string, fields []Field, file string, line int) *Entry {
	return &Entry{
		Time:   time.Now(),
		Level:  level,
		Msg:    msg,
		File:   file,
		Line:   line,
		Fields: fields,
	}
}

func (l *Logger) writeEntry(e *Entry) {
	if e.Level < l.minLevel {
		return
	}

	buf := l.bufPool.Get().(*fastBuffer)
	buf.Reset()

	l.encoder.Encode(e, buf)

	_, _ = l.out.Write(buf.Bytes())

	l.bufPool.Put(buf)
}
