package logs

import (
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func getMethod(name string) string {
	idx := strings.LastIndex(name, ".")
	if idx == -1 {
		return name
	}
	return name[idx+1:]
}

func trimExt(name string) string {
	idx := strings.LastIndex(name, ".")
	if idx == -1 {
		return name
	}
	return name[:idx]
}

const scopeKey = "scope"

// Scope adds the funcion, filename and line number to the log
func Scope(log *Logger) *Logger {
	pc, file, line, ok := runtime.Caller(1)
	if !ok {
		return log
	}
	filename := trimExt(filepath.Base(file))
	fn := runtime.FuncForPC(pc)
	if fn == nil {
		return log.With(scopeKey, filename+":"+strconv.Itoa(line))
	}
	method := getMethod(fn.Name())
	if method == "" {
		return log.With(scopeKey, filename+":"+strconv.Itoa(line))
	}
	return log.With(scopeKey, filename+"."+method+":"+strconv.Itoa(line))
}
