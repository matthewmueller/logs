package logs

import (
	"context"
	"fmt"
	"io"
	"log/slog"
	"path"
	"runtime"
	"strings"
	"sync"

	"github.com/go-logfmt/logfmt"
	"github.com/livebud/color"
)

func Console(w io.Writer) *ConsoleHandler {
	c := &ConsoleHandler{
		writer: w,
		Color:  color.Default(),
		Path:   false,
	}
	return c
}

// ConsoleHandler handler for printing logs to the terminal
type ConsoleHandler struct {
	writer io.Writer
	Color  color.Writer
	Path   bool

	// private fields
	mu     sync.Mutex // mu protects the writer
	attrs  []slog.Attr
	groups []string
}

var _ slog.Handler = (*ConsoleHandler)(nil)

// Enabled is always set to true. Use log.Filter to filter out log levels
func (c *ConsoleHandler) Enabled(context.Context, slog.Level) bool {
	return true
}

func (c *ConsoleHandler) Handle(ctx context.Context, record slog.Record) error {
	// Format the message
	msg := new(strings.Builder)
	msg.WriteString(prefix(c.Color, record.Level) + " " + record.Message)
	// Format and log the fields
	fields := new(strings.Builder)
	enc := logfmt.NewEncoder(fields)
	if record.NumAttrs() > 0 {
		prefix := strings.Join(c.groups, ".")
		record.Attrs(func(attr slog.Attr) bool {
			if prefix != "" {
				attr.Key = prefix + "." + attr.Key
			}
			enc.EncodeKeyval(attr.Key, attr.Value.String())
			return true
		})
	}
	if c.Path {
		enc.EncodeKeyval("path", caller(record.PC))
	}
	enc.Reset()
	if fields.Len() > 0 {
		msg.WriteString(" " + c.Color.Dim(fields.String()))
	}
	msg.WriteString("\n")

	// Write out
	c.mu.Lock()
	fmt.Fprint(c.writer, msg.String())
	c.mu.Unlock()

	return nil
}

func caller(pc uintptr) string {
	fs := runtime.CallersFrames([]uintptr{pc})
	f, _ := fs.Next()
	return fmt.Sprintf("%s:%d", path.Base(f.File), f.Line)
}

func (c *ConsoleHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &ConsoleHandler{
		writer: c.writer,
		Color:  c.Color,
		Path:   c.Path,
		groups: c.groups,
		attrs:  append(append([]slog.Attr{}, c.attrs...), attrs...),
	}
}

func (c *ConsoleHandler) WithGroup(group string) slog.Handler {
	return &ConsoleHandler{
		writer: c.writer,
		Color:  c.Color,
		Path:   c.Path,
		attrs:  c.attrs,
		groups: append(append([]string{}, c.groups...), group),
	}
}

var withColors = map[slog.Level]string{
	slog.LevelDebug: "|",
	slog.LevelInfo:  "|",
	slog.LevelWarn:  "|",
	slog.LevelError: "|",
}

var withoutColors = map[slog.Level]string{
	slog.LevelDebug: "debug:",
	slog.LevelInfo:  "info:",
	slog.LevelWarn:  "warn:",
	slog.LevelError: "error:",
}

// Prefixes
func prefix(color color.Writer, level slog.Level) string {
	if color.Enabled() {
		return format(color, level, withColors[level])
	}
	return withoutColors[level]
}

func format(color color.Writer, level slog.Level, msg string) string {
	switch level {
	case slog.LevelDebug:
		return color.Dim(msg)
	case slog.LevelInfo:
		return color.Blue(msg)
	case slog.LevelWarn:
		return color.Yellow(msg)
	case slog.LevelError:
		return color.Red(msg)
	default:
		return ""
	}
}
