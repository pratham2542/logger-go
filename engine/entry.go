package engine

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

func NewEntry(level LogLevel, msg string, fields []Field, file string, line int) *Entry {
	return &Entry{
		Time:   time.Now(),
		Level:  level,
		Msg:    msg,
		File:   file,
		Line:   line,
		Fields: fields,
	}
}
