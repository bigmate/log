package log

import (
	"fmt"
	"io"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logger struct {
	level zap.AtomicLevel
	base  *zap.Logger
}

func (l *logger) Debug(message string, fields ...Field) {
	l.base.Debug(message, fields...)
}

func (l *logger) Info(message string, fields ...Field) {
	l.base.Info(message, fields...)
}

func (l *logger) Warn(message string, fields ...Field) {
	l.base.Warn(message, fields...)
}

func (l *logger) Error(message string, fields ...Field) {
	l.base.Error(message, fields...)
}

func (l *logger) Fatal(message string, fields ...Field) {
	l.base.Fatal(message, fields...)
}

func (l *logger) Debugf(format string, args ...interface{}) {
	l.base.Debug(fmt.Sprintf(format, args...))
}

func (l *logger) Infof(format string, args ...interface{}) {
	l.base.Info(fmt.Sprintf(format, args...))
}

func (l *logger) Warnf(format string, args ...interface{}) {
	l.base.Warn(fmt.Sprintf(format, args...))
}

func (l *logger) Errorf(format string, args ...interface{}) {
	l.base.Error(fmt.Sprintf(format, args...))
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.base.Fatal(fmt.Sprintf(format, args...))
}

func (l *logger) With(options ...Option) Logger {
	return &logger{
		level: l.level,
		base:  l.base.WithOptions(options...),
	}
}

func (l *logger) Unwrap() *Base {
	return l.base
}

func (l *logger) Close() error {
	return l.base.Sync()
}

func encoderDefaultConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "@timestamp",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}

func defaultEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(encoderDefaultConfig())
}

const (
	callerSkipLevel    = 2
	samplingFirst      = 20
	samplingThereafter = 5
)

func newLogger(enc zapcore.Encoder, ws zapcore.WriteSyncer) *logger {
	level := zap.NewAtomicLevelAt(zapcore.InfoLevel)
	core := zapcore.NewCore(enc, ws, level)
	core = zapcore.NewSamplerWithOptions(core, time.Second, samplingFirst, samplingThereafter)
	base := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(callerSkipLevel))

	return &logger{base: base, level: level}
}

// stderr is not buffered, so there is no need to sync.
func consoleWriteSyncer() zapcore.WriteSyncer {
	return zapcore.Lock(nopSyncer{os.Stderr})
}

type nopSyncer struct {
	io.Writer
}

func (c nopSyncer) Sync() error {
	return nil
}
