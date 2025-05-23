package logs

import (
	"context"
	"log/slog"
)

// Filter logs by level
func Filter(level slog.Level, handler slog.Handler) slog.Handler {
	return &filterHandler{level, handler}
}

type filterHandler struct {
	level   slog.Level
	handler slog.Handler
}

var _ slog.Handler = (*filterHandler)(nil)

func (f *filterHandler) Enabled(ctx context.Context, l slog.Level) bool {
	return l >= f.level
}

func (f *filterHandler) Handle(ctx context.Context, record slog.Record) error {
	return f.handler.Handle(ctx, record)
}

func (f *filterHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &filterHandler{
		level:   f.level,
		handler: f.handler.WithAttrs(attrs),
	}
}

func (f *filterHandler) WithGroup(group string) slog.Handler {
	return &filterHandler{
		level:   f.level,
		handler: f.handler.WithGroup(group),
	}
}
