package logs

import (
	"context"
	"log/slog"
)

// Filter logs by level
func Filter(level Level, handler slog.Handler) slog.Handler {
	return &LevelHandler{level, handler}
}

type LevelHandler struct {
	level   slog.Level
	handler slog.Handler
}

var _ slog.Handler = (*LevelHandler)(nil)

func (f *LevelHandler) Enabled(ctx context.Context, l slog.Level) bool {
	return l >= f.level
}

func (f *LevelHandler) Handle(ctx context.Context, record slog.Record) error {
	return f.handler.Handle(ctx, record)
}

func (f *LevelHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &LevelHandler{f.level, f.handler.WithAttrs(attrs)}
}

func (f *LevelHandler) WithGroup(group string) slog.Handler {
	return &LevelHandler{f.level, f.handler.WithGroup(group)}
}
