package logger

import "go.uber.org/zap/zapcore"

// Field type
type Field = zapcore.Field

// Debug level information
func Debug(msg string, fields ...Field) {
	getLogger().Debug(msg, fields...)
}

// Info level information
func Info(msg string, fields ...Field) {
	getLogger().Info(msg, fields...)
}

// Warn level information
func Warn(msg string, fields ...Field) {
	getLogger().Warn(msg, fields...)
}

// Error level information
func Error(msg string, fields ...Field) {
	getLogger().Error(msg, fields...)
}

// Panic level information
func Panic(msg string, fields ...Field) {
	getLogger().Panic(msg, fields...)
}

// InfoPf format level information
func InfoPf(format string, a ...interface{}) {
	getSugaredLogger().Infof(format, a...)
}
