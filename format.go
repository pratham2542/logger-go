package logger

import (
	"os"
	"runtime"
)

func (l *Logger) log(level LogLevel, msg string, fields ...Field) {
	file, line := l.captureCallerIfEnabled()
	e := newEntry(level, msg, fields, file, line)
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

// Public API
func (l *Logger) Debug(msg string, fields ...Field) { l.log(DEBUG, msg, fields...) }
func (l *Logger) Info(msg string, fields ...Field)  { l.log(INFO, msg, fields...) }
func (l *Logger) Warn(msg string, fields ...Field)  { l.log(WARN, msg, fields...) }
func (l *Logger) Error(msg string, fields ...Field) { l.log(ERROR, msg, fields...) }
func (l *Logger) Fatal(msg string, fields ...Field) {
	l.log(FATAL, msg, fields...)
	os.Exit(1)
}
