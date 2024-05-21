package utils

import (
	"context"
	"io"
	"log/slog"
)

type CustomJsonLogger struct {
	logger *slog.Logger
}

func NewCustomJsonLogger(w io.Writer, logLevel slog.Leveler) CustomJsonLogger {
	opts := slog.HandlerOptions{
		Level: logLevel,
	}
	logger := slog.New(slog.NewJSONHandler(w, &opts))

	return CustomJsonLogger{logger: logger}
}

func (l *CustomJsonLogger) Info(ctx context.Context, message string, args ...any) {
	requestId := getRequestId(ctx)
	args = append(args, string(RequestIdName), requestId)

	l.logger.Info(message, args...)
}

func (l *CustomJsonLogger) Error(ctx context.Context, message string, args ...any) {
	requestId := getRequestId(ctx)
	args = append(args, string(RequestIdName), requestId)

	l.logger.Error(message, args...)
}

func (l *CustomJsonLogger) Debug(ctx context.Context, message string, args ...any) {
	requestId := getRequestId(ctx)
	args = append(args, string(RequestIdName), requestId)

	l.logger.Debug(message, args...)
}

func (l *CustomJsonLogger) Warn(ctx context.Context, message string, args ...any) {
	requestId := getRequestId(ctx)
	args = append(args, string(RequestIdName), requestId)

	l.logger.Warn(message, args...)
}

func getRequestId(ctx context.Context) string {
	requestId := ctx.Value(RequestIdName)
	if requestId == nil || requestId == "" {
		requestId = "application"
	}

	return requestId.(string)
}
