package consumer

import (
	"context"
	"fmt"
	"sync"
)

// ConsumerManager — управление несколькими инстансами
type ConsumerManager struct {
	consumers map[string]*Consumer
	mu        sync.RWMutex
	wg        sync.WaitGroup
}

func NewManager() *ConsumerManager {
	return &ConsumerManager{
		consumers: make(map[string]*Consumer),
	}
}

func (m *ConsumerManager) Add(c *Consumer) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if _, exists := m.consumers[c.id]; exists {
		return fmt.Errorf("consumer with ID %s already exists", c.id)
	}
	m.consumers[c.id] = c
	return nil
}

func (m *ConsumerManager) StartAll(ctx context.Context) error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	for _, c := range m.consumers {
		c := c
		m.wg.Add(1)
		go func() {
			defer m.wg.Done()
			topics := c.topics
			if err := c.Start(ctx, topics); err != nil {
				c.logger.Printf("Start error: %v", err)
			}
		}()
	}
	return nil
}

func (m *ConsumerManager) StopAll() error {
	m.mu.RLock()
	defer m.mu.RUnlock()

	var errs []error
	for _, c := range m.consumers {
		if err := c.Stop(); err != nil {
			errs = append(errs, err)
		}
	}
	m.wg.Wait()

	if len(errs) > 0 {
		return fmt.Errorf("shutdown errors: %v", errs)
	}
	return nil
}

func (m *ConsumerManager) Get(id string) (*Consumer, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	c, ok := m.consumers[id]
	return c, ok
}
