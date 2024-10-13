package logger

import (
	"context"
	"log/slog"
	"os"
)

type Port interface {
	Info(msg string, atrs ...any)
	Debug(msg string, atrs ...any)
	Warn(msg string, atrs ...any)
	Error(msg string, atrs ...any)

	InfoContext(ctx context.Context, msg string, atrs ...any)
	DebugContext(ctx context.Context, msg string, atrs ...any)
	WarnContext(ctx context.Context, msg string, atrs ...any)
	ErrorContext(ctx context.Context, msg string, atrs ...any)
}

type Adapter struct {
	*slog.Logger
}

func New() *Adapter {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	return &Adapter{logger}
}
