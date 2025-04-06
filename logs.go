package logs

import (
	"log/slog"
	"os"
)

type Handler = slog.Handler
type Logger = slog.Logger
type Record = slog.Record

func New(handler slog.Handler) *Logger {
	return slog.New(handler)
}

func Default() *Logger {
	handler := slog.Handler(Console(os.Stderr))
	level := LevelInfo
	if lvl, ok := parseLevel(os.Getenv("LOG")); ok {
		level = lvl
	}
	if debug, ok := os.LookupEnv("DEBUG"); ok {
		handler = packageHandler(handler, debug)
	} else {
		handler = Filter(level, handler)
	}
	return slog.New(handler)
}

// Fatal calls the default logger and exits with status code 1. Only intended
// for use in main functions or top-level handlers.
func Fatal(err error) {
	Default().Error(err.Error())
	os.Exit(1)
}
