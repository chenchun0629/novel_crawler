package log

import (
	"fmt"
	"os"
)

var defaultHelper *Helper

// Debug logs a message at debug level.
func Debug(a ...interface{}) {
	defaultHelper.Log(LevelDebug, fmt.Sprint(a...))
}

// Debugf logs a message at debug level.
func Debugf(format string, a ...interface{}) {
	defaultHelper.Log(LevelDebug, fmt.Sprintf(format, a...))
}

// Debugw logs a message at debug level.
func Debugw(msg string, keyvals ...interface{}) {
	defaultHelper.Logw(LevelDebug, msg, keyvals...)
}

// Info logs a message at info level.
func Info(a ...interface{}) {
	defaultHelper.Log(LevelInfo, fmt.Sprint(a...))
}

// Infof logs a message at info level.
func Infof(format string, a ...interface{}) {
	defaultHelper.Log(LevelInfo, fmt.Sprintf(format, a...))
}

// Infow logs a message at info level.
func Infow(msg string, keyvals ...interface{}) {
	defaultHelper.Logw(LevelInfo, msg, keyvals...)
}

// Warn logs a message at warn level.
func Warn(a ...interface{}) {
	defaultHelper.Log(LevelWarn, fmt.Sprint(a...))
}

// Warnf logs a message at warnf level.
func Warnf(format string, a ...interface{}) {
	defaultHelper.Log(LevelWarn, fmt.Sprintf(format, a...))
}

// Warnw logs a message at warnf level.
func Warnw(msg string, keyvals ...interface{}) {
	defaultHelper.Logw(LevelWarn, msg, keyvals...)
}

// Error logs a message at error level.
func Error(a ...interface{}) {
	defaultHelper.Log(LevelError, fmt.Sprint(a...))
}

// Errorf logs a message at error level.
func Errorf(format string, a ...interface{}) {
	defaultHelper.Log(LevelError, fmt.Sprintf(format, a...))
}

// Errorw logs a message at error level.
func Errorw(msg string, keyvals ...interface{}) {
	defaultHelper.Logw(LevelError, msg, keyvals...)
}

// Fatal logs a message at fatal level.
func Fatal(a ...interface{}) {
	defaultHelper.Log(LevelFatal, fmt.Sprint(a...))
	os.Exit(1)
}

// Fatalf logs a message at fatal level.
func Fatalf(format string, a ...interface{}) {
	defaultHelper.Log(LevelFatal, fmt.Sprintf(format, a...))
	os.Exit(1)
}

// Fatalw logs a message at fatal level.
func Fatalw(msg string, keyvals ...interface{}) {
	defaultHelper.Logw(LevelFatal, msg, keyvals...)
	os.Exit(1)
}
