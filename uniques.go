package logs

import "log/slog"

func uniqueAttrs(list []slog.Attr) []slog.Attr {
	seen := map[string]struct{}{}
	set := make([]slog.Attr, 0, len(list))
	for i := len(list) - 1; i >= 0; i-- {
		item := list[i]
		if _, ok := seen[item.Key]; ok {
			continue
		}
		seen[item.Key] = struct{}{}
		set = append(set, item)
	}
	return set
}
