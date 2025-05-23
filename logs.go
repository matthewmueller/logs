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

// Default logger writes to stderr at the info level by default. This can be
// adjusted by setting the LOG environment variable to one of the levels:
// debug, info, warn or error.
func Default() *Logger {
	level := LevelInfo
	if lvl, ok := parseLevel(os.Getenv("LOG")); ok {
		level = lvl
	}
	return slog.New(Filter(level, Console(os.Stderr)))
}

// Fatal calls the default logger and exits with status code 1. Only intended
// for use in main functions or top-level handlers.
func Fatal(err error) {
	Default().Error(err.Error())
	os.Exit(1)
}
