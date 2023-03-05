package dbutils

import (
	"context"

	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/vmedinskiy/highload/internal/pkg/log"
)

type Logger struct {
	logger *zap.Logger
}

func NewLogger() *Logger {
	return &Logger{logger: log.L().Desugar().WithOptions(zap.AddCallerSkip(1))}
}

func (dl *Logger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	fields := make([]zapcore.Field, len(data))
	i := 0
	for k, v := range data {
		fields[i] = zap.Any(k, v)
		i++
	}
	// ll := dl.logger.With(zap.String("reqId", requestid.FromContext(ctx)))
	ll := dl.logger
	switch level {
	case pgx.LogLevelTrace:
		ll.Debug(msg, append(fields, zap.Stringer("PGX_LOG_LEVEL", level))...)
	case pgx.LogLevelDebug:
		ll.Debug(msg, fields...)
	case pgx.LogLevelInfo:
		ll.Info(msg, fields...)
	case pgx.LogLevelWarn:
		ll.Warn(msg, fields...)
	case pgx.LogLevelError:
		ll.Error(msg, fields...)
	default:
		ll.Error(msg, append(fields, zap.Stringer("PGX_LOG_LEVEL", level))...)
	}
}
