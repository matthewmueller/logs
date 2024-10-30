# logs

[![Go Reference](https://pkg.go.dev/badge/github.com/matthewmueller/logs.svg)](https://pkg.go.dev/github.com/matthewmueller/logs)

Log handlers for the new [log/slog](https://go.dev/blog/slog) package.

## Features

- Built on top of the new [log/slog](https://go.dev/blog/slog) package
- Pretty `console` handler for terminals
- Adds a level filter handler
- Adds a multi-logger

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

// Integrates well with libraries that use the standard *slog.Logger
func ThirdPartyFunc(log *slog.Logger) {
  // ...
}

// Works! log is still a *slog.Logger
ThirdPartyFunc(log)
```

## Contributors

- Matt Mueller ([@mattmueller](https://twitter.com/mattmueller))

## License

MIT
