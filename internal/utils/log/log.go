package log

import (
	"context"

	"go.uber.org/zap"
)

var (
	logger *zap.Logger
)

func Init() *zap.Logger {
	logger, _ := zap.NewDevelopment()
	return logger
}

func GetLogger(ctx context.Context) *zap.Logger {
	if logger == nil {
		return Init()
	}
	return logger
}

func Error(ctx context.Context, msg string) {
	GetLogger(ctx).Error(msg)
}

func Fatal(ctx context.Context, msg string) {
	GetLogger(ctx).Fatal(msg)
}

func Info(ctx context.Context, msg string) {
	GetLogger(ctx).Info(msg)
}
