package kafkaconsumer

import (
	"context"
	"sync"
	"time"
)

type Deduper interface {
	Seen(id string) bool // уже видели (и обработали)
	MarkDone(id string)  // помечаем как обработанное
}

type inMemoryDeduper struct {
	mu   sync.RWMutex
	data map[string]time.Time
	ttl  time.Duration
}

func NewInMemoryDeduper(ctx context.Context, ttl time.Duration) *inMemoryDeduper {
	d := &inMemoryDeduper{
		data: make(map[string]time.Time, 1024),
		ttl:  ttl,
	}

	go d.runGC(ctx)
	return d
}

func (d *inMemoryDeduper) Seen(id string) bool {
	d.mu.RLock()
	_, ok := d.data[id]
	d.mu.RUnlock()
	return ok
}

func (d *inMemoryDeduper) MarkDone(id string) {
	d.mu.Lock()
	d.data[id] = time.Now()
	d.mu.Unlock()
}

func (d *inMemoryDeduper) runGC(ctx context.Context) {
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
