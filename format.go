package logger

import (
	"bytes"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

func (l *Logger) log(level LogLevel, msg string, args ...interface{}) {
	if level < l.minLevel {
		return
	}

	buf := l.bufPool.Get().(*bytes.Buffer) // Type assertion will always succeed so no check required
	buf.Reset()                            // clears any old data in the buffer pool

	// Timestamp
	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	buf.WriteString("[")
	buf.WriteString(timestamp)
	buf.WriteString("] [")

	// Level
	buf.WriteString(levelNames[level])
	buf.WriteString("] [")

	// Caller
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		file = "???"
		line = 0
	}

	if l.fullPath {
		buf.WriteString(file)
	} else {
		fileParts := strings.Split(file, "/")
		buf.WriteString(fileParts[len(fileParts)-1])
	}

	buf.WriteString(":")
	buf.WriteString(strconv.Itoa(line))
	buf.WriteString("] ")

	writeArgs(buf, msg, args...)

	buf.WriteByte('\n')

	l.out.Write(buf.Bytes())

	l.bufPool.Put(buf) // return the buffer to the pool for next time use
}

// Public API
func (l *Logger) Debug(msg string, args ...interface{}) { l.log(DEBUG, msg, args...) }
func (l *Logger) Info(msg string, args ...interface{})  { l.log(INFO, msg, args...) }
func (l *Logger) Warn(msg string, args ...interface{})  { l.log(WARN, msg, args...) }
func (l *Logger) Error(msg string, args ...interface{}) { l.log(ERROR, msg, args...) }
func (l *Logger) Fatal(msg string, args ...interface{}) {
	l.log(FATAL, msg, args...)
	os.Exit(1)
}

func writeArgs(buf *bytes.Buffer, msg string, args ...interface{}) {
	buf.WriteString(msg)
	for _, arg := range args {
		buf.WriteByte(' ')
		switch v := arg.(type) {
		case string:
			buf.WriteString(v)
		case int:
			buf.WriteString(strconv.Itoa(v))
		case int64:
			buf.WriteString(strconv.FormatInt(v, 10))
		case uint:
			buf.WriteString(strconv.FormatUint(uint64(v), 10))
		case uint64:
			buf.WriteString(strconv.FormatUint(v, 10))
		case float32:
			buf.WriteString(strconv.FormatFloat(float64(v), 'f', -1, 32))
		case float64:
			buf.WriteString(strconv.FormatFloat(v, 'f', -1, 64))
		case bool:
			buf.WriteString(strconv.FormatBool(v))
		case error:
			buf.WriteString(v.Error())
		default:
			// fallback – only here we call fmt, rarely
			buf.WriteString(fmt.Sprint(v))
		}
	}
}
