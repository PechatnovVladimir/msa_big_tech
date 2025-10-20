package inbox

import (
	"time"
)

type searchMessageOptions struct {
	NotBefore *time.Time
	NotAfter  *time.Time

	//AggregateType *AggregateType
	//EventType     *EventType

	OnlyUnpublished bool
	MaxRetryCount   int
	DueAt           *time.Time
	Limit           int
	WithLock        bool
}

type SearchMessageOption func(o *searchMessageOptions)

func WithLimit(n int) SearchMessageOption {
	return func(o *searchMessageOptions) { o.Limit = n }
}

func WithOnlyUnpublished() SearchMessageOption {
	return func(o *searchMessageOptions) { o.OnlyUnpublished = true }
}

func WithMaxRetryCount(n int) SearchMessageOption {
	return func(o *searchMessageOptions) { o.MaxRetryCount = n }
}

func WithLock() SearchMessageOption {
	return func(o *searchMessageOptions) { o.WithLock = true }
}

//func WithAggregateType(t AggregateType) SearchMessageOption {
//	return func(o *searchMessageOptions) { o.AggregateType = &t }
//}
//
//func WithEventType(t EventType) SearchMessageOption {
//	return func(o *searchMessageOptions) { o.EventType = &t }
//}

func WithNotBefore(t time.Time) SearchMessageOption {
	return func(o *searchMessageOptions) { o.NotBefore = &t }
}

func WithNotAfter(t time.Time) SearchMessageOption {
	return func(o *searchMessageOptions) { o.NotAfter = &t }
}

func WithDueAt(t time.Time) SearchMessageOption {
	return func(o *searchMessageOptions) { o.DueAt = &t }
}

func CollectSearchMessageOptions(opts ...SearchMessageOption) searchMessageOptions {
	res := searchMessageOptions{
		Limit:         10,
		MaxRetryCount: 3,
	}

	for _, opt := range opts {
		opt(&res)
	}
	return res
}

type updateMessageOptions struct {
	// window по времени для partition pruning (created_at)
	NotBefore *time.Time
	NotAfter  *time.Time

	//AggregateType *AggregateType
	//EventType     *EventType

	IDs []string

	// что обновляем
	SetPublishedAt   *time.Time
	IncRetryBy       int
	SetNextAttemptAt *time.Time

	// фильтры статуса
	OnlyUnpublished bool // по умолчанию true
}

type UpdateMessageOption func(*updateMessageOptions)

func CollectUpdateMessageOptions(opts ...UpdateMessageOption) updateMessageOptions {
	o := updateMessageOptions{
		OnlyUnpublished: true,
	}
	for _, opt := range opts {
		opt(&o)
	}
	return o
}

// ----- Option builders -----

func WithUpdateIDs(ids ...string) UpdateMessageOption {
	return func(o *updateMessageOptions) { o.IDs = append(o.IDs, ids...) }
}

//func WithUpdateAggregateType(at AggregateType) UpdateMessageOption {
//	return func(o *updateMessageOptions) { o.AggregateType = &at }
//}
//
//func WithUpdateEventType(et EventType) UpdateMessageOption {
//	return func(o *updateMessageOptions) { o.EventType = &et }
//}

func WithUpdateNotBefore(t time.Time) UpdateMessageOption {
	return func(o *updateMessageOptions) { o.NotBefore = &t }
}

func WithUpdateNotAfter(t time.Time) UpdateMessageOption {
	return func(o *updateMessageOptions) { o.NotAfter = &t }
}

func SetPublishedAt(ts time.Time) UpdateMessageOption {
	return func(o *updateMessageOptions) { o.SetPublishedAt = &ts }
}

func IncRetry(by int) UpdateMessageOption {
	return func(o *updateMessageOptions) { o.IncRetryBy = by }
}

func IncludePublished() UpdateMessageOption {
	return func(o *updateMessageOptions) { o.OnlyUnpublished = false }
}

func SetNextAttemptAt(ts time.Time) UpdateMessageOption {
	return func(o *updateMessageOptions) { o.SetNextAttemptAt = &ts }
}
