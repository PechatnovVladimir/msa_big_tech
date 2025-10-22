package inbox

import (
	"time"
)

type searchMessageOptions struct {
	Status      []string
	MaxAttempts int
	Limit       int
	WithLock    bool
}

type SearchMessageOption func(o *searchMessageOptions)

func WithLimit(n int) SearchMessageOption {
	return func(o *searchMessageOptions) { o.Limit = n }
}

func WithStatus(status []string) SearchMessageOption {
	return func(o *searchMessageOptions) { o.Status = status }
}

func WithMaxAttempts(n int) SearchMessageOption {
	return func(o *searchMessageOptions) { o.MaxAttempts = n }
}

func WithLock() SearchMessageOption {
	return func(o *searchMessageOptions) { o.WithLock = true }
}

func CollectSearchMessageOptions(opts ...SearchMessageOption) searchMessageOptions {
	res := searchMessageOptions{
		Limit:       10,
		MaxAttempts: 10,
	}

	for _, opt := range opts {
		opt(&res)
	}
	return res
}

type updateMessageOptions struct {
	IDs         []string
	Status      string
	LastError   string
	IncAttempts int
	ProcessedAt time.Time
}

type UpdateMessageOption func(*updateMessageOptions)

func CollectUpdateMessageOptions(opts ...UpdateMessageOption) updateMessageOptions {
	o := updateMessageOptions{}
	for _, opt := range opts {
		opt(&o)
	}
	return o
}

func WithUpdateIDs(ids ...string) UpdateMessageOption {
	return func(o *updateMessageOptions) { o.IDs = append(o.IDs, ids...) }
}

func WithUpdateStatus(status string) UpdateMessageOption {
	return func(o *updateMessageOptions) { o.Status = status }
}

func WithUpdateAttempts(n int) UpdateMessageOption {
	return func(o *updateMessageOptions) { o.IncAttempts = 1 }
}

func WithUpdateLastError(lastError string) UpdateMessageOption {
	return func(o *updateMessageOptions) { o.LastError = lastError }
}

func WithUpdateProcessedAt(processedAt time.Time) UpdateMessageOption {
	return func(o *updateMessageOptions) { o.ProcessedAt = processedAt }
}
