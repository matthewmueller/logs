package logs

import (
	"log/slog"
	"os"
)

type Handler = slog.Handler
type Logger = slog.Logger

func Default() *Logger {
	console := Console(os.Stderr)
	return slog.New(console)
}

func New(handler slog.Handler) *Logger {
	return slog.New(handler)
}
