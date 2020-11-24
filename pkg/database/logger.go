package database

import (
	"context"

	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const skipCallerLogger = 9

type loggerWrapZap2Pgx struct {
	log *zap.SugaredLogger
}

func getLevelFromZap(log *zap.SugaredLogger) pgx.LogLevel {
	var levels = map[pgx.LogLevel]zapcore.Level{
		pgx.LogLevelTrace: zapcore.DebugLevel,
		pgx.LogLevelDebug: zapcore.DebugLevel,
		pgx.LogLevelInfo:  zapcore.InfoLevel,
		pgx.LogLevelWarn:  zapcore.WarnLevel,
		pgx.LogLevelError: zapcore.ErrorLevel,
	}

	core := log.Desugar().Core()
	prevLevel := pgx.LogLevelError
	for level := pgx.LogLevelError; level <= pgx.LogLevelTrace; level++ {
		if core.Enabled(levels[pgx.LogLevel(level)]) {
			prevLevel = level
		} else {
			return pgx.LogLevel(prevLevel)
		}
	}
	return pgx.LogLevelTrace
}

func newLoggerWrapPgx(log *zap.SugaredLogger) *loggerWrapZap2Pgx {
	log = log.Desugar().
		WithOptions(zap.AddCallerSkip(skipCallerLogger)).
		Sugar()
	return &loggerWrapZap2Pgx{
		log: log,
	}
}

func (l *loggerWrapZap2Pgx) Log(_ context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	fields := make([]interface{}, 0, len(data)*2)
	for name, value := range data {
		fields = append(fields, name, value)
	}
	switch level {
	case pgx.LogLevelTrace, pgx.LogLevelDebug:
		l.log.Debugw(msg, fields...)
	case pgx.LogLevelInfo:
		l.log.Infow(msg, fields...)
	case pgx.LogLevelWarn:
		l.log.Warnw(msg, fields...)
	case pgx.LogLevelError:
		l.log.Errorw(msg, fields...)
	}
}
