package logger_test

import (
	"io"
	"testing"

	"github.com/pratham2542/logger-go"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// --- Custom logger benchmark ---
func BenchmarkCustomLogger(b *testing.B) {
	l := logger.NewLogger(logger.DEBUG, io.Discard, false)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Info("Benchmarking custom logger", "i", i, "ok", true, "f", 3.1415)
	}
}

// --- Zap logger benchmark ---
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
		zapcore.NewConsoleEncoder(encoderCfg),
		zapcore.AddSync(io.Discard),
		zap.DebugLevel,
	)
	l := zap.New(core)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		l.Info("Benchmarking zap logger",
			zap.Int("i", i),
			zap.Bool("ok", true),
			zap.Float64("f", 3.1415),
		)
	}
}
