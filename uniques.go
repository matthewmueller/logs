package logs

import "log/slog"

func uniqueAttrs(list []slog.Attr) []slog.Attr {
	seen := map[string]struct{}{}
	set := make([]slog.Attr, 0, len(list))
	for _, item := range list {
		if _, ok := seen[item.Key]; ok {
			continue
		}
		seen[item.Key] = struct{}{}
		set = append(set, item)
	}
	return set
}
