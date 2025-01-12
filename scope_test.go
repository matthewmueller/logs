package logs_test

import (
	"bytes"
	"log/slog"
	"strings"
	"testing"

	"github.com/livebud/color"
	"github.com/matryer/is"
	"github.com/matthewmueller/logs"
)

func TestScope(t *testing.T) {
	is := is.New(t)
	buf := new(bytes.Buffer)
	log := logs.New(slog.NewTextHandler(buf, nil))
	log = logs.Scope(log)
	log.Info("hello world")
	actual := buf.String()
	// Line number based on what line `log = logs.Scope(log)` is called
	is.True(strings.Contains(actual, "scope=scope_test.TestScope:18"))
}

func another(log *slog.Logger) {
	logs.Scope(log).Info("hello world")
}

func TestScopeMultiple(t *testing.T) {
	is := is.New(t)
	buf := new(bytes.Buffer)
	log := logs.New(slog.NewTextHandler(buf, &slog.HandlerOptions{AddSource: true}))
	log = logs.Scope(log)
	log.WithGroup("a").WithGroup("b").With("key", "value").Info("hello world")
	another(log)
	actual := buf.String()
	// Line number based on what line `log = logs.Scope(log)` is called
	is.True(strings.Contains(actual, "scope=scope_test.TestScopeMultiple:33"))
	is.True(strings.Contains(actual, "scope=scope_test.another:26"))
}

func TestScopeOverride(t *testing.T) {
	is := is.New(t)
	buf := new(bytes.Buffer)
	console := logs.Console(buf)
	console.Color = color.Ignore()
	log := logs.New(console)
	log = logs.Scope(log)
	log = logs.Scope(log)
	log.Info("hello world")
	actual := buf.String()
	// Line number based on what line `log = logs.Scope(log)` is called
	is.True(!strings.Contains(actual, "scope=scope_test.TestScopeOverride:48"))
	is.True(strings.Contains(actual, "scope=scope_test.TestScopeOverride:49"))
}

func BenchmarkScope(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buf := new(bytes.Buffer)
		log := logs.New(slog.NewTextHandler(buf, nil))
		log = logs.Scope(log)
		log.Info("hello world")
	}
}
