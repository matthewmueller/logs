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
		Source: false,
	}
	return c
}

// ConsoleHandler handler for printing logs to the terminal
type ConsoleHandler struct {
	Color  color.Writer
	Source bool

	// mu protects the writer
	mu     sync.Mutex
	writer io.Writer

	// private fields
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
	prefix := strings.Join(c.groups, ".")
	if record.NumAttrs() > 0 {
		record.Attrs(func(attr slog.Attr) bool {
			key := attr.Key
			if prefix != "" {
				key = prefix + "." + key
			}
			enc.EncodeKeyval(key, attr.Value.String())
			return true
		})
	}
	for _, attr := range c.attrs {
		key := attr.Key
		if prefix != "" {
			key = prefix + "." + key
		}
		enc.EncodeKeyval(key, attr.Value.String())
	}
	if c.Source {
		enc.EncodeKeyval(slog.SourceKey, caller(record.PC))
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
		Source: c.Source,
		groups: c.groups,
		// Prioritize the new attributes over the old ones
		attrs: uniqueAttrs(append(c.attrs, attrs...)),
	}
}

func (c *ConsoleHandler) WithGroup(group string) slog.Handler {
	return &ConsoleHandler{
		writer: c.writer,
		Color:  c.Color,
		Source: c.Source,
		groups: append(c.groups, group),
		attrs:  c.attrs,
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
