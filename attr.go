package logs

import (
	"log/slog"
)

type Attr = slog.Attr

// Aliases for slog package types
var (
	String   = slog.String
	Int      = slog.Int
	Int64    = slog.Int64
	Uint64   = slog.Uint64
	Float64  = slog.Float64
	Bool     = slog.Bool
	Any      = slog.Any
	Time     = slog.Time
	Duration = slog.Duration
)

// Err creates an error attribute for logging.
func Err(err error) Attr {
	return Attr{
		Key:   "err",
		Value: slog.StringValue(err.Error()),
	}
}
