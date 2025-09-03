package logger

import (
	"strconv"
	"sync/atomic"
	"time"
)

type timestampCache struct {
	val atomic.Value // stores []byte
}

func newTimestampCache() *timestampCache {
	tc := &timestampCache{}
	tc.update() // initialize once

	// Background refresher
	go func() {
		ticker := time.NewTicker(time.Millisecond) // update every 1ms
		defer ticker.Stop()

		for t := range ticker.C {
			tc.set(t)
		}
	}()
	return tc
}

func (tc *timestampCache) update() {
	tc.set(time.Now())
}

func (tc *timestampCache) set(t time.Time) {
	// Prebuild timestamp format: YYYY-MM-DD HH:MM:SS.mmm
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	millis := t.Nanosecond() / 1e6

	// Allocate once per refresh
	buf := make([]byte, 0, 24)
	buf = append(buf, '[')
	buf = appendInt(buf, year, 4)
	buf = append(buf, '-')
	buf = appendInt(buf, int(month), 2)
	buf = append(buf, '-')
	buf = appendInt(buf, day, 2)
	buf = append(buf, ' ')
	buf = appendInt(buf, hour, 2)
	buf = append(buf, ':')
	buf = appendInt(buf, min, 2)
	buf = append(buf, ':')
	buf = appendInt(buf, sec, 2)
	buf = append(buf, '.')
	buf = appendInt(buf, millis, 3)
	buf = append(buf, ']')

	tc.val.Store(buf)
}

func (tc *timestampCache) Bytes() []byte {
	return tc.val.Load().([]byte)
}

// appendInt ensures zero-padded width
func appendInt(buf []byte, n, width int) []byte {
	s := strconv.Itoa(n)
	for len(s) < width {
		s = "0" + s
	}
	return append(buf, s...)
}
