package logger

import (
	"context"

	"go.uber.org/zap"
)

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}
func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Ctx(ctx context.Context) *zap.Logger {
	traceID := ctx.Value("trace_id")
	if traceID != nil {
		return logger.With(zap.String("trace_id", traceID.(string)))
	}
	return logger
}
