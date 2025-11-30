package logger_test

import (
	"io"
	"testing"

	"github.com/pratham2542/logger-go"
	eng "github.com/pratham2542/logger-go/engine"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func BenchmarkCustomLogger(b *testing.B) {
	engine := eng.NewEngine(eng.DEBUG, eng.DefaultTextEncoder(), io.Discard)
	l := logger.NewLogger(engine)

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for range 100 {
				l.Info("Benchmarking custom logger",
					eng.Int("i", 42),
					eng.Bool("ok", true),
					eng.Float("f", 3.1415),
				)
			}
		}
	})
}

func BenchmarkZapLogger(b *testing.B) {
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

	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			for range 100 {
				l.Info("Benchmarking zap logger",
					zap.Int("i", 42),
					zap.Bool("ok", true),
					zap.Float64("f", 3.1415),
				)
			}
		}
	})
}
