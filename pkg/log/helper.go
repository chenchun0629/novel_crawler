package log

import (
	"context"
	"fmt"
	"os"
)

var DefaultMessageKey = "msg"

// Option is Helper option.
type Option func(*Helper)

// Helper is a logger helper.
type Helper struct {
	logger Logger
	msgKey string
}

func WithMessageKey(k string) Option {
	return func(opts *Helper) {
		opts.msgKey = k
	}
}

// NewHelper new a logger helper.
func NewHelper(logger Logger, opts ...Option) *Helper {
	options := &Helper{
		msgKey: DefaultMessageKey, // default message key
		logger: logger,
	}
	for _, o := range opts {
		o(options)
	}
	return options
}

// WithContext returns a shallow copy of h with its context changed
// to ctx. The provided ctx must be non-nil.
func (h *Helper) WithContext(ctx context.Context) *Helper {
	return &Helper{
		logger: WithContext(ctx, h.logger),
	}
}

// Log Print log by level and keyvals.
func (h *Helper) Log(level Level, keyvals ...interface{}) {
	_ = h.logger.Log(level, keyvals...)
}

// Log Print log by level and keyvals.
func (h *Helper) Logw(level Level, msg string, keyvals ...interface{}) {
	_ = h.logger.Logw(level, msg, keyvals...)
}

// Debug logs a message at debug level.
func (h *Helper) Debug(a ...interface{}) {
	h.Log(LevelDebug, fmt.Sprint(a...))
}

// Debugf logs a message at debug level.
func (h *Helper) Debugf(format string, a ...interface{}) {
	h.Log(LevelDebug, fmt.Sprintf(format, a...))
}

// Debugw logs a message at debug level.
func (h *Helper) Debugw(msg string, keyvals ...interface{}) {
	h.Logw(LevelDebug, msg, keyvals...)
}

// Info logs a message at info level.
func (h *Helper) Info(a ...interface{}) {
	h.Log(LevelInfo, fmt.Sprint(a...))
}

// Infof logs a message at info level.
func (h *Helper) Infof(format string, a ...interface{}) {
	h.Log(LevelInfo, fmt.Sprintf(format, a...))
}

// Infow logs a message at info level.
func (h *Helper) Infow(msg string, keyvals ...interface{}) {
	h.Logw(LevelInfo, msg, keyvals...)
}

// Warn logs a message at warn level.
func (h *Helper) Warn(a ...interface{}) {
	h.Log(LevelWarn, fmt.Sprint(a...))
}

// Warnf logs a message at warnf level.
func (h *Helper) Warnf(format string, a ...interface{}) {
	h.Log(LevelWarn, fmt.Sprintf(format, a...))
}

// Warnw logs a message at warnf level.
func (h *Helper) Warnw(msg string, keyvals ...interface{}) {
	h.Logw(LevelWarn, msg, keyvals...)
}

// Error logs a message at error level.
func (h *Helper) Error(a ...interface{}) {
	h.Log(LevelError, fmt.Sprint(a...))
}

// Errorf logs a message at error level.
func (h *Helper) Errorf(format string, a ...interface{}) {
	h.Log(LevelError, fmt.Sprintf(format, a...))
}

// Errorw logs a message at error level.
func (h *Helper) Errorw(msg string, keyvals ...interface{}) {
	h.Logw(LevelError, msg, keyvals...)
}

// Fatal logs a message at fatal level.
func (h *Helper) Fatal(a ...interface{}) {
	h.Log(LevelFatal, fmt.Sprint(a...))
	os.Exit(1)
}

// Fatalf logs a message at fatal level.
func (h *Helper) Fatalf(format string, a ...interface{}) {
	h.Log(LevelFatal, fmt.Sprintf(format, a...))
	os.Exit(1)
}

// Fatalw logs a message at fatal level.
func (h *Helper) Fatalw(msg string, keyvals ...interface{}) {
	h.Logw(LevelFatal, msg, keyvals...)
	os.Exit(1)
}
