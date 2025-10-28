package consumer

import (
	"context"
	"sync"
	"time"
)

type InMemoryDeduplicator struct {
	mu   sync.RWMutex
	data map[string]time.Time
	ttl  time.Duration
}

func NewInMemory(ctx context.Context, ttl time.Duration) *InMemoryDeduplicator {
	d := &InMemoryDeduplicator{
		data: make(map[string]time.Time, 1024),
		ttl:  ttl,
	}
	go d.runGC(ctx)
	return d
}

func (d *InMemoryDeduplicator) Seen(id string) bool {
	d.mu.RLock()
	_, ok := d.data[id]
	d.mu.RUnlock()
	return ok
}

func (d *InMemoryDeduplicator) MarkSeen(id string) {
	d.mu.Lock()
	d.data[id] = time.Now()
	d.mu.Unlock()
}

func (d *InMemoryDeduplicator) runGC(ctx context.Context) {
	t := time.NewTicker(d.ttl / 2)
	defer t.Stop()
	for {
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			cut := time.Now().Add(-d.ttl)
			d.mu.Lock()
			for k, ts := range d.data {
				if ts.Before(cut) {
					delete(d.data, k)
				}
			}
			d.mu.Unlock()
		}
	}
}
