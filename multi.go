package logs

import (
	"context"

	"log/slog"

	"golang.org/x/sync/errgroup"
)

func Multi(handlers ...slog.Handler) MultiHandler {
	return MultiHandler(handlers)
}

type MultiHandler []slog.Handler

var _ slog.Handler = (MultiHandler)(nil)

func (handlers MultiHandler) Enabled(ctx context.Context, lvl slog.Level) bool {
	return true
}

func (handlers MultiHandler) Handle(ctx context.Context, record slog.Record) error {
	eg, ctx := errgroup.WithContext(ctx)
	for _, handler := range handlers {
		handler := handler
		if !handler.Enabled(ctx, record.Level) {
			continue
		}
		eg.Go(func() error { return handler.Handle(ctx, record.Clone()) })
	}
	if err := eg.Wait(); err != nil {
		return err
	}
	return nil
}

func (handlers MultiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	clone := make(MultiHandler, len(handlers))
	for i, handler := range handlers {
		clone[i] = handler.WithAttrs(attrs)
	}
	return clone
}

func (handlers MultiHandler) WithGroup(group string) slog.Handler {
	clone := make(MultiHandler, len(handlers))
	for i, handler := range handlers {
		clone[i] = handler.WithGroup(group)
	}
	return clone
}
