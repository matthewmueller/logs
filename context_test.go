package logs_test

import (
	"context"
	"testing"

	"github.com/matryer/is"
	"github.com/matthewmueller/logs"
)

func TestWithLogAndFromContext(t *testing.T) {
	is := is.New(t)

	// Create a new context
	ctx := context.Background()

	// Create a new Logger
	logger := logs.Default()

	// Add the Logger to the context
	ctxWithLog := logs.WithLog(ctx, logger)

	// Retrieve the Logger from the context
	retrievedLogger := logs.FromContext(ctxWithLog)

	// Assert that the retrieved Logger is the same as the one added
	is.True(retrievedLogger == logger) // Logger should match

	// Test retrieving a Logger from a context without a Logger
	emptyLogger := logs.FromContext(ctx)
	is.True(emptyLogger == nil) // Should return nil
}
