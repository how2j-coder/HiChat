package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"strings"
)

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

// Sync flushing any buffered log entries, applications should take care to call Sync before exiting.
func Sync() error {
	_ = getSugaredLogger().Sync()
	err := getLogger().Sync()
	if err != nil && !strings.Contains(err.Error(), "/dev/stdout") {
		return err
	}
	return nil
}

// Err type
func Err(err error) Field {
	return zap.Error(err)
}

// Any type, if it is a composite type such as object, slice, map, etc., use Any
func Any(key string, val interface{}) Field {
	return zap.Any(key, val)
}

// String type
func String(key string, val string) Field {
	return zap.String(key, val)
}
