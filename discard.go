package logs

import (
	"context"
	"log/slog"
)

type discard struct{}

func (discard) Enabled(context.Context, slog.Level) bool  { return false }
func (discard) Handle(context.Context, slog.Record) error { return nil }
func (discard) WithAttrs([]slog.Attr) slog.Handler        { return discard{} }
func (discard) WithGroup(string) slog.Handler             { return discard{} }

func Discard() slog.Handler {
	return discard{}
}
