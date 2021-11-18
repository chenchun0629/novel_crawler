package log

import (
	"go.uber.org/zap"
)

var _ Logger = (*zapLogger)(nil)

type zapLogger struct {
	log *zap.Logger
}

// NewZapLogger new a logger with writer.
func NewZapLogger() Logger {
	return &zapLogger{
		log: zap.NewExample(),
	}
}

// Log print the kv pairs log.
func (l *zapLogger) Log(level Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 {
		return nil
	}
	if (len(keyvals) & 1) == 1 {
		keyvals = append(keyvals, "KEYVALS UNPAIRED")
	}

	switch level {
	case LevelDebug:
		l.log.Sugar().Debug(keyvals...)
	case LevelInfo:
		l.log.Sugar().Info(keyvals...)
	case LevelWarn:
		l.log.Sugar().Warn(keyvals...)
	case LevelError:
		l.log.Sugar().Error(keyvals...)
	case LevelFatal:
		l.log.Sugar().Fatal(keyvals...)
	default:
		panic("invalid log level")
	}

	return nil
}

func (l *zapLogger) Logw(level Level, msg string, keyvals ...interface{}) error {
	if (len(keyvals) & 1) == 1 {
		keyvals = append(keyvals, "KEYVALS UNPAIRED")
	}

	switch level {
	case LevelDebug:
		l.log.Sugar().Debugw(msg, keyvals...)
	case LevelInfo:
		l.log.Sugar().Infow(msg, keyvals...)
	case LevelWarn:
		l.log.Sugar().Warnw(msg, keyvals...)
	case LevelError:
		l.log.Sugar().Errorw(msg, keyvals...)
	case LevelFatal:
		l.log.Sugar().Fatalw(msg, keyvals...)
	default:
		panic("invalid log level")
	}

	return nil
}
