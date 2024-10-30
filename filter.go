package logs

import (
	"context"
	"log/slog"
)

// Filter logs by level
func Filter(level slog.Level, handler slog.Handler) *FilterHandler {
	return &FilterHandler{level, handler}
}

type FilterHandler struct {
	level   slog.Level
	handler slog.Handler
}

var _ slog.Handler = (*FilterHandler)(nil)

func (f *FilterHandler) Enabled(ctx context.Context, l slog.Level) bool {
	return l >= f.level
}

func (f *FilterHandler) Handle(ctx context.Context, record slog.Record) error {
	return f.handler.Handle(ctx, record)
}

func (f *FilterHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &FilterHandler{
		level:   f.level,
		handler: f.handler.WithAttrs(attrs),
	}
}

func (f *FilterHandler) WithGroup(group string) slog.Handler {
	return &FilterHandler{
		level:   f.level,
		handler: f.handler.WithGroup(group),
	}
}
