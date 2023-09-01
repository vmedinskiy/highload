package log

import (
	"context"
	"fmt"

	"github.com/go-faster/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type contextLoggerKey struct{}

type Logger struct {
	*zap.SugaredLogger
}

var globalSugaredLogger *Logger

func Init(level uint) error {
	var config zap.Config
	if level == 0 {
		config = zap.NewProductionConfig()
	} else {
		if level > 2 {
			level = 2
		}
		config = zap.NewDevelopmentConfig()
	}

	config.Level.SetLevel([]zapcore.Level{zap.PanicLevel, zap.InfoLevel, zap.DebugLevel}[level])
	config.Encoding = "json"

	lg, err := config.Build()
	if err != nil {
		return errors.Wrap(err, "building logger from config")
	}
	globalSugaredLogger = &Logger{lg.Sugar()}
	return nil
}

func L() *Logger {
	return globalSugaredLogger
}

func WithLogger(ctx context.Context, logger *Logger) context.Context {
	return context.WithValue(ctx, contextLoggerKey{}, logger)
}

// C returns logger, taken from context.
func C(ctx context.Context) *Logger {
	logger, ok := ctx.Value(contextLoggerKey{}).(*Logger)
	if ok {
		return logger
	}

	if globalSugaredLogger != nil {
		return globalSugaredLogger
	}

	return &Logger{zap.L().Sugar()}
}

func (l *Logger) With(args ...any) *Logger {
	return &Logger{l.SugaredLogger.With(args...)}
}

func Sync() {
	err := globalSugaredLogger.Sync()
	if err != nil {
		fmt.Println(err.Error())
	}
}
