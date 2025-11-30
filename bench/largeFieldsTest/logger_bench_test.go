package logger_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/pratham2542/logger-go"
	eng "github.com/pratham2542/logger-go/engine"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func BenchmarkCustomLogger70Fields(b *testing.B) {
	engine := eng.NewEngine(eng.DEBUG, eng.DefaultTextEncoder(), io.Discard)
	l := logger.NewLogger(engine)

	fields := make([]eng.Field, 0, 70000)
	for i := 0; i < 70000; i++ {
		fields = append(fields, eng.Int(fmt.Sprintf("k%d", i), i))
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for range 10 {
				l.Info("Benchmarking 70 fields", fields...)
			}
		}
	})
}

func BenchmarkZapLogger70Fields(b *testing.B) {
	encoderCfg := zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderCfg),
		zapcore.AddSync(io.Discard),
		zap.DebugLevel,
	)
	l := zap.New(core)

	fields := make([]zap.Field, 0, 70000)
	for i := 0; i < 70000; i++ {
		fields = append(fields, zap.Int(fmt.Sprintf("k%d", i), i))
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for range 10 {
				l.Info("Benchmarking 70 fields", fields...)
			}
		}
	})
}
