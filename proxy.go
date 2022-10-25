package log

import (
	"strings"

	"go.uber.org/zap"
)

var globalLogger = newLogger(defaultEncoder(), consoleWriteSyncer())

// Logger is an abstraction over various log destinations.
type Logger interface {
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})

	Debug(message string, fields ...Field)
	Info(message string, fields ...Field)
	Warn(message string, fields ...Field)
	Error(message string, fields ...Field)
	Fatal(message string, fields ...Field)

	With(options ...Option) Logger
	Unwrap() *Base
}

func Debugf(format string, args ...interface{}) {
	globalLogger.Debugf(format, args...)
}

func Infof(format string, args ...interface{}) {
	globalLogger.Infof(format, args...)
}

func Warnf(format string, args ...interface{}) {
	globalLogger.Warnf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	globalLogger.Errorf(format, args...)
}

func Fatalf(format string, args ...interface{}) {
	globalLogger.Fatalf(format, args...)
}

func Debug(message string, field ...Field) {
	globalLogger.Debug(message, field...)
}

func Info(message string, fields ...Field) {
	globalLogger.Info(message, fields...)
}

func Warn(message string, fields ...Field) {
	globalLogger.Warn(message, fields...)
}

func Error(message string, fields ...Field) {
	globalLogger.Error(message, fields...)
}

func Fatal(message string, fields ...Field) {
	globalLogger.Fatal(message, fields...)
}

func With(options ...Option) Logger {
	return globalLogger.With(options...)
}

// Base will get deprecated, it's for compatibility only.
type Base = zap.Logger

// Unwrap  will get deprecated, it's for compatibility only.
func Unwrap() *Base {
	return globalLogger.base
}

func Global() Logger {
	return globalLogger
}

func SetLevel(level Level) {
	globalLogger.level.SetLevel(level)
}

// Close does not make any effect as stderr is not
// buffered and is default destination for logs.
func Close() error {
	return globalLogger.Close()
}

func LevelFromString(level string) Level {
	switch strings.ToLower(level) {
	case "debug":
		return DebugLevel
	case "info":
		return InfoLevel
	case "warn":
		return WarnLevel
	case "fatal":
		return FatalLevel
	default:
		return ErrorLevel
	}
}
