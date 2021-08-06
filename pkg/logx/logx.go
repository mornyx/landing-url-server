// Package logx simply wraps the zap logger for global use.
package logx

import (
	"go.uber.org/zap"
)

var lg *zap.Logger

func init() {
	var err error
	lg, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}
}

/* TODO: make logger configurable. */

func Logger() *zap.Logger {
	return lg
}

func Debug(msg string, fields ...zap.Field) {
	lg.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	lg.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	lg.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	lg.Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	lg.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	lg.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	lg.Fatal(msg, fields...)
}
