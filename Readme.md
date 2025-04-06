# logs

[![Go Reference](https://pkg.go.dev/badge/github.com/matthewmueller/logs.svg)](https://pkg.go.dev/github.com/matthewmueller/logs)

Simple, pretty logger for your Go apps.

```go
log := logs.Default()
log.WithGroup("grouped").Debug("debug line", "path", "console_test.go")
log.Info("some info")
log.Warn("some warning")
log.Error("an error", "err", errors.New("oh no"))
```

![log screenshot](https://github.com/user-attachments/assets/8b108829-8dbc-4a50-86c5-94363d1dca7b)

## Features

- Compatible with [slog](https://go.dev/blog/slog)
- Pretty `console` handler for terminals
- Adds a level filter handler
- Adds a concurrent multi-logger

## Install

```sh
go get github.com/matthewmueller/logs
```

## Example

```go
// Go with the default logger
log := logs.Default()

// Or configure
log = logs.New(
  logs.Multi(
    logs.Filter(slog.LevelInfo, logs.Console(os.Stderr)),
    slog.NewJSONHandler(os.Stderr, nil),
  ),
)

log.WithGroup("hello").Debug("world", "args", 10)
log.Info("hello", "planet", "world", "args", 10)
log.Warn("hello", "planet", "world", "args", 10)
log.Error("hello world", "planet", "world", "args", 10)

// Integrates well with other libraries because log is still a *slog.Logger
var logger *slog.Logger = log
logger.WithGroup("hello").Debug("world", "args", 10)
logger.Info("hello", "planet", "world", "args", 10)
logger.Warn("hello", "planet", "world", "args", 10)
logger.Error("hello world", logs.String("planet", "world"), "args", 10)
```

## Contributors

- Matt Mueller ([@mattmueller](https://twitter.com/mattmueller))

## License

MIT
