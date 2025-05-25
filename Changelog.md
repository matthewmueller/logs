# 0.0.8 / 2025-05-25

- **BREAKING**: switch back to `Discard()` returning a `*slog.Logger`

# 0.0.7 / 2025-05-23

- **BREAKING**: `logs.Default()` now logs at the info level
- **Breaking**: `logs.Discard()` now returns a `slog.Handler` instead of a `*slog.Logger`.
- **Breaking**: removed `logs.Info`, `logs.Error`, `logs.Debug` and `logs.Warn`
- **Fixed**: duplicate attrs bug
- **Added**: `slog.MustParseLevel`, `slog.Err(err)` attribute, `slog.Fatal(err)`

# 0.0.6 / 2025-01-12

- add global default logger methods:
  - Debug, DebugContext
  - Info, InfoContext
  - Warn, WarnContext
  - Error, ErrorContext

# 0.0.5 / 2025-01-12

- fix: support `logger.With(...)` attrs in console handler
- new: dedupe attrs passed in `logger.With(...)` in console handler
- new: add `logs.Scope(logger)` key for annotating functions with `scope=<file>.<function>:<line>`

# 0.0.4 / 2025-01-12

- add discard logger
- remove path attr by default
- replace fluent api

# 0.0.3 / 2024-11-03

- add a license

# 0.0.2 / 2024-11-02

- don't add space between message and attrs when no attrs

# 0.0.1 / 2024-10-30

- initial commit
