package logs

import (
	"fmt"
	"log/slog"
	"strings"
)

type Level = slog.Level

var (
	// Aliases for slog package levels
	LevelDebug = slog.LevelDebug
	LevelInfo  = slog.LevelInfo
	LevelWarn  = slog.LevelWarn
	LevelError = slog.LevelError
)

// ParseLevel parses a string into a log level
func ParseLevel(level string) (Level, error) {
	lvl, ok := parseLevel(level)
	if !ok {
		return 0, fmt.Errorf("logs: %q is not a valid level", level)
	}
	return lvl, nil
}

// MustParseLevel parses a string into a log level and panics if it fails
func MustParseLevel(level string) Level {
	lvl, err := ParseLevel(level)
	if err != nil {
		panic(err)
	}
	return lvl
}

// ParseLevel parses a string into a log level
func parseLevel(field string) (Level, bool) {
	switch strings.ToLower(field) {
	case "debug":
		return slog.LevelDebug, true
	case "info":
		return slog.LevelInfo, true
	case "warn":
		return slog.LevelWarn, true
	case "error":
		return slog.LevelError, true
	default:
		return 0, false
	}
}
