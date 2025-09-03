package logger

import (
	"encoding"
	"fmt"
	"os"
	"path"
	"runtime"
	"strconv"
)

func (l *Logger) log(level LogLevel, msg string, args ...any) {
	if level < l.minLevel {
		return
	}

	buf := l.bufPool.Get().(*fastBuffer) // Type assertion will always succeed so no check required
	buf.Reset()                          // clears any old data in the buffer pool

	// Timestamp
	buf.Write(l.tsCache.Bytes())
	buf.WriteString(" [")

	// Level
	buf.WriteString(levelNames[level])
	buf.WriteString("] ")

	if l.withCaller {
		_, file, line, ok := runtime.Caller(2)
		if !ok {
			file = "???"
			line = 0
		}
		buf.WriteString("[")

		if l.fullPath {
			buf.WriteString(file)
		} else {
			// Load from the cache to reduce finding the base filename
			if val, ok := l.fileCache.Load(file); ok {
				buf.WriteString(val.(string))
			} else {
				base := path.Base(file)
				l.fileCache.Store(file, base)
				buf.WriteString(base)
			}
		}

		buf.WriteString(":")
		buf.WriteString(strconv.Itoa(line))
		buf.WriteString("] ")
	}

	writeArgs(buf, msg, args...)

	buf.WriteByte('\n')
	l.out.Write(buf.Bytes())

	buf.TrimAndPut(&l.bufPool) // return the buffer to the pool for next time use (Trimed for larger buffer)
}

// Public API
func (l *Logger) Debug(msg string, args ...any) { l.log(DEBUG, msg, args...) }
func (l *Logger) Info(msg string, args ...any)  { l.log(INFO, msg, args...) }
func (l *Logger) Warn(msg string, args ...any)  { l.log(WARN, msg, args...) }
func (l *Logger) Error(msg string, args ...any) { l.log(ERROR, msg, args...) }
func (l *Logger) Fatal(msg string, args ...any) {
	l.log(FATAL, msg, args...)
	os.Exit(1)
}

func writeArgs(buf *fastBuffer, msg string, args ...any) {
	buf.WriteString(msg)
	for _, arg := range args {
		buf.WriteByte(' ')
		switch v := arg.(type) {
		case string:
			buf.WriteString(v)
		case int:
			buf.b = strconv.AppendInt(buf.b, int64(v), 10)
		case int64:
			buf.b = strconv.AppendInt(buf.b, v, 10)
		case uint:
			buf.b = strconv.AppendUint(buf.b, uint64(v), 10)
		case uint64:
			buf.b = strconv.AppendUint(buf.b, v, 10)
		case float32:
			buf.b = strconv.AppendFloat(buf.b, float64(v), 'f', -1, 32)
		case float64:
			buf.b = strconv.AppendFloat(buf.b, v, 'f', -1, 64)
		case bool:
			buf.b = strconv.AppendBool(buf.b, v)
		case error:
			buf.WriteString(v.Error())

		// added custome types to use its own String or toString method before defaulting to fmt for string conversion
		case fmt.Stringer:
			buf.WriteString(v.String())
		case encoding.TextMarshaler:
			if b, err := v.MarshalText(); err == nil {
				buf.b = append(buf.b, b...)
			} else {
				buf.WriteString(fmt.Sprint(v)) // fallback if MarshalText fails
			}

		default:
			// fallback – only here we call fmt, rarely
			buf.WriteString(fmt.Sprint(v))
		}
	}
}
