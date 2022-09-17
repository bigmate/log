package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type (
	Level  = zapcore.Level
	Field  = zapcore.Field
	Option = zap.Option
)

const (
	DebugLevel = zapcore.DebugLevel
	InfoLevel  = zapcore.InfoLevel
	WarnLevel  = zapcore.WarnLevel
	ErrorLevel = zapcore.ErrorLevel
	FatalLevel = zapcore.FatalLevel
)

var (
	Binary    = zap.Binary
	Bool      = zap.Bool
	Float64   = zap.Float64
	Float32   = zap.Float32
	Int       = zap.Int
	Int64     = zap.Int64
	Int32     = zap.Int32
	Int16     = zap.Int16
	Int8      = zap.Int8
	String    = zap.String
	Uint      = zap.Uint
	Uint64    = zap.Uint64
	Uint32    = zap.Uint32
	Uint16    = zap.Uint16
	Uint8     = zap.Uint8
	Namespace = zap.Namespace
	Stringer  = zap.Stringer
	Time      = zap.Time
	Duration  = zap.Duration
	Reflect   = zap.Reflect
	Err       = zap.Error
	Any       = zap.Any

	Fields        = zap.Fields
	AddCallerSkip = zap.AddCallerSkip
)
