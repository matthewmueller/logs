package logs_test

import (
	"testing"

	"log/slog"

	"github.com/matryer/is"
	"github.com/matthewmueller/logs"
)

func TestIncludeHandler(t *testing.T) {
	is := is.New(t)

	// Create a recorder to capture log records
	rec := logs.Recorder()

	// Define include functions
	includeInfo := func(rec slog.Record) bool {
		return rec.Level == slog.LevelInfo
	}
	includeError := func(rec slog.Record) bool {
		return rec.Level == slog.LevelError
	}

	// Create an Include handler with the recorder and include functions
	handler := logs.Include(rec, includeInfo, includeError)

	// Create a logger with the handler
	logger := slog.New(handler)

	// Log messages with different levels
	logger.Info("This is an info message")
	logger.Debug("This is a debug message")
	logger.Error("This is an error message")

	// Retrieve the recorded logs
	records := rec.Records()

	// Verify that only the included levels are logged
	is.Equal(len(records), 2) // expected 2 records

	is.True(records[0].Level == slog.LevelInfo) // first record should be info level
	is.Equal(records[0].Message, "This is an info message")

	is.True(records[1].Level == slog.LevelError) // second record should be error level
	is.Equal(records[1].Message, "This is an error message")
}
