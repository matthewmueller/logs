package logs

import (
	"fmt"
	"log/slog"
)

// ParseLevel parses a string into a log level
func ParseLevel(level string) (slog.Level, error) {
	switch level {
	case "debug":
		return slog.LevelDebug, nil
	case "info":
		return slog.LevelInfo, nil
	case "warn":
		return slog.LevelWarn, nil
	case "error":
		return slog.LevelError, nil
	}
	return 0, fmt.Errorf("logs: %q is not a valid level", level)
}
