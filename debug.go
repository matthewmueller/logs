package logs

import (
	"log/slog"
	"os"
)

// Debugger returns a logger that writes to stderr with a level of Debug.
func Debugger() *Logger {
	return slog.New(Console(os.Stderr))
}
