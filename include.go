package logs

import (
	"context"
	"log/slog"
)

// Include filters logs based on a custom function.
func Include(handler Handler, includes ...func(rec slog.Record) bool) Handler {
	return &includeHandler{handler, includes}
}

type includeHandler struct {
	handler  Handler
	includes []func(rec slog.Record) bool
}

var _ slog.Handler = (*includeHandler)(nil)

func (f *includeHandler) Enabled(ctx context.Context, l slog.Level) bool {
	return true
}

func (f *includeHandler) include(record slog.Record) bool {
	for _, fn := range f.includes {
		if fn(record) {
			return true
		}
	}
	return false
}

func (f *includeHandler) Handle(ctx context.Context, record slog.Record) error {
	if !f.include(record) {
		return nil
	}
	return f.handler.Handle(ctx, record)
}

func (f *includeHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &includeHandler{f.handler.WithAttrs(attrs), f.includes}
}

func (f *includeHandler) WithGroup(group string) slog.Handler {
	return &includeHandler{f.handler.WithGroup(group), f.includes}
}
