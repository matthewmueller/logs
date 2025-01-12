package logs_test

import (
	"bytes"
	"errors"
	"log/slog"
	"os"
	"strings"
	"testing"
	"testing/slogtest"

	"github.com/livebud/color"
	"github.com/matryer/is"
	"github.com/matthewmueller/logs"
)

func TestConsoleHandler(t *testing.T) {
	t.Skip("skip until: https://github.com/golang/go/issues/61706")
	is := is.New(t)
	buf := new(bytes.Buffer)
	console := logs.Console(buf)
	console.Color = color.Ignore()
	err := slogtest.TestHandler(console, func() []map[string]any {
		lines := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")
		results := make([]map[string]any, len(lines))
		for i, line := range lines {
			result := map[string]any{}
			result[slog.LevelKey] = nil
			result[slog.TimeKey] = nil
			line = strings.TrimLeft(line, "info: ")
			pairs := strings.Split(line, " ")
			result[slog.MessageKey] = pairs[0]
			for _, pair := range pairs[1:] {
				kv := strings.SplitN(pair, "=", 2)
				if len(kv) != 2 {
					continue
				}
				result[kv[0]] = kv[1]
			}
			results[i] = result
		}
		return results
	})
	is.NoErr(err)
}

func TestConsole(t *testing.T) {
	is := is.New(t)
	buf := new(bytes.Buffer)
	console := logs.Console(buf)
	log := logs.New(console)
	log.WithGroup("group").Debug("world", "args", 10)
	log.Info("hello", "planet", "world", "args", 10)
	log.Warn("hello", "planet", "world", "args", 10)
	log.Error("hello world", "planet", "world", "args", 10)
	lines := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")
	is.Equal(len(lines), 4)
}

func TestConsoleWith(t *testing.T) {
	is := is.New(t)
	buf := new(bytes.Buffer)
	console := logs.Console(buf)
	console.Color = color.Ignore()
	log := logs.New(console)
	log.With("planet", "world").Info("hello", "args", 10, "args", 20)
	log.With("planet", "world").Warn("hello", "args", 10)
	log.With("planet", "world").Error("hello", "args", 10)
	lines := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")
	is.Equal(len(lines), 3)
	is.Equal(string(lines[0]), "info: hello args=10 args=20 planet=world")
	is.Equal(string(lines[1]), "warn: hello args=10 planet=world")
	is.Equal(string(lines[2]), "error: hello args=10 planet=world")
}

func TestConsoleDedupe(t *testing.T) {
	is := is.New(t)
	buf := new(bytes.Buffer)
	console := logs.Console(buf)
	console.Color = color.Ignore()
	log := logs.New(console)
	log.With("planet", "earth").With("planet", "mars").Info("hello", "args", 10, "args", 20)
	lines := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")
	is.Equal(len(lines), 1)
	// Only attributes added via With are deduped.
	is.Equal(string(lines[0]), "info: hello args=10 args=20 planet=mars")
}

func TestGroup(t *testing.T) {
	is := is.New(t)
	buf := new(bytes.Buffer)
	console := logs.Console(buf)
	console.Color = color.Ignore()
	log := logs.New(console)
	l1 := log.WithGroup("one")
	l1.Debug("world", "args", 10)
	l2 := log.WithGroup("two").WithGroup("three")
	l2.Debug("hello", "args", 10)
	l1.Info("hello", "amount", 20)
	lines := strings.Split(strings.TrimRight(buf.String(), "\n"), "\n")
	is.Equal(len(lines), 3)
	is.Equal(string(lines[0]), "debug: world one.args=10")
	is.Equal(string(lines[1]), "debug: hello two.three.args=10")
	is.Equal(string(lines[2]), "info: hello one.amount=20")
}

func ExampleConsole() {
	console := logs.Console(os.Stdout)
	console.Color = color.Ignore()
	log := logs.New(console)
	log.WithGroup("grouped").Debug("debug line", "path", "console_test.go")
	log.Info("some info")
	log.Warn("some warning")
	log.Error("an error", "err", errors.New("oh no"))
	// Output:
	// debug: debug line grouped.path=console_test.go
	// info: some info
	// warn: some warning
	// error: an error err="oh no"
}

func ExampleDiscard() {
	log := logs.Discard()
	var logger *slog.Logger = log
	logger.WithGroup("hello").Debug("world", "args", 10)
	logger.Info("hello", "planet", "world", "args", 10)
	logger.Warn("hello", "planet", "world", "args", 10)
	logger.Error("hello world", slog.String("planet", "world"), "args", 10)
	// Output:
}

func ExampleLogger() {
	console := logs.Console(os.Stdout)
	console.Color = color.Ignore()
	log := logs.New(console)
	var logger *slog.Logger = log
	logger.WithGroup("hello").Debug("world", "args", 10)
	logger.Info("hello", "planet", "world", "args", 10)
	logger.Warn("hello", "planet", "world", "args", 10)
	logger.Error("hello world", slog.String("planet", "world"), "args", 10)
	// Output:
	// debug: world hello.args=10
	// info: hello planet=world args=10
	// warn: hello planet=world args=10
	// error: hello world planet=world args=10
}
