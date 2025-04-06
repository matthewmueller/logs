package logs

import "context"

type logKey struct{}

// WithLog adds a Logger to the context.
func WithLog(ctx context.Context, log *Logger) context.Context {
	return context.WithValue(ctx, logKey{}, log)
}

// FromContext retrieves the Logger from the context.
func FromContext(ctx context.Context) *Logger {
	if log, ok := ctx.Value(logKey{}).(*Logger); ok {
		return log
	}
	return nil
}
