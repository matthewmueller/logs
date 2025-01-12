package logs_test

import (
	"bytes"
	"log/slog"
	"testing"

	"github.com/matryer/is"
	"github.com/matthewmueller/logs"
)

func TestFilterDebug(t *testing.T) {
	t.Setenv("NO_COLOR", "1")
	is := is.New(t)
	buf := new(bytes.Buffer)
	console := logs.Console(buf)
	log := slog.New(logs.Filter(slog.LevelInfo, console))
	log.Debug("hello", "args", 10)
	log.Info("hello", "planet", "world", "args", 10)
	log.Warn("hello", "planet", "world", "args", 10)
	log.Error("hello world", "planet", "world", "args", 10)
	lines := bytes.Split(bytes.TrimRight(buf.Bytes(), "\n"), []byte("\n"))
	is.Equal(len(lines), 3)
	is.Equal(string(lines[0]), "info: hello planet=world args=10")
	is.Equal(string(lines[1]), "warn: hello planet=world args=10")
	is.Equal(string(lines[2]), "error: hello world planet=world args=10")
}

func TestFilterError(t *testing.T) {
	t.Setenv("NO_COLOR", "1")
	is := is.New(t)
	buf := new(bytes.Buffer)
	console := logs.Console(buf)
	log := slog.New(logs.Filter(slog.LevelError, console))
	log.Debug("hello", "args", 10)
	log.Info("hello", "planet", "world", "args", 10)
	log.Warn("hello", "planet", "world", "args", 10)
	log.Error("hello world", "planet", "world", "args", 10)
	lines := bytes.Split(bytes.TrimRight(buf.Bytes(), "\n"), []byte("\n"))
	is.Equal(len(lines), 1)
	is.Equal(string(lines[0]), "error: hello world planet=world args=10")
}
