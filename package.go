package logs

import (
	"log/slog"
	"runtime"
	"strings"
)

func packageHandler(handler slog.Handler, pattern string) slog.Handler {
	if pattern == "" {
		return handler
	}
	return Include(handler, compilePattern(pattern))
}

func compilePattern(pattern string) func(rec slog.Record) bool {
	fields := strings.Fields(pattern)
	includes := make([]func(path string) bool, len(fields))
	for i, field := range strings.Fields(pattern) {
		includes[i] = func(path string) bool {
			return strings.HasPrefix(path, field)
		}
	}
	return func(rec slog.Record) bool {
		pkgPath := packagePath(rec.PC)
		for _, fn := range includes {
			if fn(pkgPath) {
				return true
			}
		}
		return false
	}
}

func packagePath(pc uintptr) string {
	fs := runtime.CallersFrames([]uintptr{pc})
	f, _ := fs.Next()
	if f.File == "" {
		return ""
	}

	// Strip GOPATH or GOROOT prefix if present
	parts := strings.Split(f.Function, "/")
	if len(parts) > 0 {
		fn := parts[len(parts)-1]
		if i := strings.LastIndex(fn, "."); i != -1 {
			return fn[:i] // package path before function
		}
	}

	return ""
}
