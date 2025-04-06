package logs_test

import (
	"fmt"
	"log/slog"
	"testing"

	"github.com/matryer/is"
	"github.com/matthewmueller/logs"
)

func TestRecorder(t *testing.T) {
	is := is.New(t)
	rec := logs.Recorder()
	log := logs.New(rec)
	log.Info("hello")
	log.Error("world", logs.String("foo", "bar"))
	log.With("foo", "bar").Debug("baz", logs.String("zap", "zoo"))
	log.WithGroup("foo").Warn("bar")
	records := rec.Records()

	// Ensure the correct number of records
	is.Equal(len(records), 4)

	// Check the first record
	first := records[0]
	is.Equal(first.Level, slog.LevelInfo)
	is.Equal(first.Message, "hello")
	is.Equal(first.NumAttrs(), 0)

	// Check the second record
	second := records[1]
	is.Equal(second.Level, slog.LevelError)
	is.Equal(second.Message, "world")
	is.Equal(second.NumAttrs(), 1)
	var secondAttrs []slog.Attr
	second.Attrs(func(a slog.Attr) bool {
		secondAttrs = append(secondAttrs, a)
		return true
	})
	is.Equal(len(secondAttrs), 1)
	is.Equal(secondAttrs[0].Key, "foo")
	is.Equal(secondAttrs[0].Value.String(), "bar")

	// Check the third record
	third := records[2]
	is.Equal(third.Level, slog.LevelDebug)
	is.Equal(third.Message, "baz")
	is.Equal(third.NumAttrs(), 2)
	var thirdAttrs []slog.Attr
	third.Attrs(func(a slog.Attr) bool {
		fmt.Println(a.Key, a.Value)
		thirdAttrs = append(thirdAttrs, a)
		return true
	})
	is.Equal(len(thirdAttrs), 2)
	is.Equal(thirdAttrs[0].Key, "zap")
	is.Equal(thirdAttrs[0].Value.String(), "zoo")
	is.Equal(thirdAttrs[1].Key, "foo")
	is.Equal(thirdAttrs[1].Value.String(), "bar")

	// Check the fourth record
	fourth := records[3]
	is.Equal(fourth.Level, slog.LevelWarn)
	is.Equal(fourth.Message, "bar")
	is.Equal(fourth.NumAttrs(), 0)
}
