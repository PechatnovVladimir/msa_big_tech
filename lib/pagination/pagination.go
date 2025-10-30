package pagination

import "time"

var (
	ASC  = false
	DESC = true
)

type SortField struct {
	Name string
	Desc bool
}

type Cursor struct {
	ID   *string
	Time *time.Time
}

type Option func(opts *Options)

type Options struct {
	limit  uint64
	cursor Cursor
	sort   []SortField
}

func NewOptions(opts ...Option) Options {
	var o Options
	for _, opt := range opts {
		opt(&o)
	}
	return o
}

func (p *Options) Limit() uint64 {
	return p.limit
}

func (p *Options) Cursor() Cursor {
	return p.cursor
}

func (p *Options) OrderBy() []SortField {
	return p.sort
}

func OrderBy(name string, desc bool) SortField {
	return SortField{Name: name, Desc: desc}
}

func WithLimit(limit uint64) Option {
	return func(opts *Options) { opts.limit = limit }
}

func WithCursor(cursor Cursor) Option {
	return func(opts *Options) { opts.cursor = cursor }
}

func WithSortFields(fields ...SortField) Option {
	return func(opts *Options) { opts.sort = fields }
}
