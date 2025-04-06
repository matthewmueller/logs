package logs

import (
	"context"
	"log/slog"
	"strings"
)

func Recorder() *Rec {
	return &Rec{
		records: &[]slog.Record{},
	}
}

type Rec struct {
	records *[]slog.Record
	groups  []string
	attrs   []slog.Attr
}

var _ slog.Handler = (*Rec)(nil)

func (r *Rec) Records() []slog.Record {
	return *r.records
}

func (r *Rec) Enabled(context.Context, slog.Level) bool {
	return true
}

func (r *Rec) Handle(ctx context.Context, record slog.Record) error {
	prefix := strings.Join(r.groups, ".")
	for _, attr := range r.attrs {
		key := attr.Key
		if prefix != "" {
			key = prefix + "." + key
		}
		record.AddAttrs(slog.Attr{
			Key:   key,
			Value: attr.Value,
		})
	}
	*r.records = append(*r.records, record)
	return nil
}

func (r *Rec) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &Rec{
		records: r.records,
		groups:  r.groups,
		attrs:   uniqueAttrs(append(r.attrs, attrs...)),
	}
}

func (r *Rec) WithGroup(group string) slog.Handler {
	return &Rec{
		records: r.records,
		groups:  append(r.groups, group),
		attrs:   r.attrs,
	}
}
