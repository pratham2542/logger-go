package logger_test

import (
	"fmt"
	"io"
	"testing"

	"github.com/pratham2542/logger-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func BenchmarkCustomLogger70Fields(b *testing.B) {
	l := logger.NewLogger(logger.DEBUG, io.Discard, false, false)

	fields := make([]any, 0, 140)
	for i := 0; i < 70; i++ {
		fields = append(fields, fmt.Sprintf("k%d", i), i)
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Info("Benchmarking 70 fields", fields...)
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

	fields := make([]zap.Field, 0, 70)
	for i := 0; i < 70; i++ {
		fields = append(fields, zap.Int(fmt.Sprintf("k%d", i), i))
	}

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			l.Info("Benchmarking 70 fields", fields...)
		}
	})
}
