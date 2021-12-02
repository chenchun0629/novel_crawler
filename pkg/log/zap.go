package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var _ Logger = (*zapLogger)(nil)

type zapLogger struct {
	log *zap.Logger
}

func (l *zapLogger) Defer() {
	_ = l.log.Sync()
}

// NewZapLogger new a logger with writer.
func NewZapLogger() Logger {
	var (
		encoderCfg = zapcore.EncoderConfig{
			MessageKey:     "msg",
			LevelKey:       "level",
			TimeKey:        "time",
			NameKey:        "logger",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.999999999"),
			EncodeDuration: zapcore.StringDurationEncoder,
		}
		core = zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), os.Stdout, zap.DebugLevel)
		zl   = zap.New(core).WithOptions()
	)

	return &zapLogger{
		log: zl,
	}
}

// Log print the kv pairs log.
func (l *zapLogger) Log(level Level, keyvals ...interface{}) error {
	if len(keyvals) == 0 {
		return nil
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
