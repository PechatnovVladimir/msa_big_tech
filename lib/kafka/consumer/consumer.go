package consumer

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"os"
	"sync"
	"time"
)

type ConsumerHandler interface {
	HandleBatch(context.Context, []*sarama.ConsumerMessage) error
}

type Consumer struct {
	topics           []string
	consumerGroup    sarama.ConsumerGroup
	handler          ConsumerHandler
	deduplicator     Deduplicator
	ready            chan bool
	wg               sync.WaitGroup
	id               string
	batchSize        int
	batchTimeout     time.Duration
	flushOnRebalance bool
	logger           *log.Logger
}

func NewConsumer(cfg Config, handler ConsumerHandler, dedup Deduplicator) (*Consumer, error) {
	if cfg.ConsumerID == "" {
		return nil, fmt.Errorf("ConsumerID is required")
	}

	if cfg.SaramaConfig == nil {
		cfg.SaramaConfig = sarama.NewConfig()
		cfg.SaramaConfig.Version = sarama.V2_5_0_0
		cfg.SaramaConfig.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRange()
		cfg.SaramaConfig.Consumer.Offsets.Initial = sarama.OffsetOldest
		cfg.SaramaConfig.Consumer.Return.Errors = true // ОБЯЗАТЕЛЬНО ЧИТАЕМ cg.group.Errors()

	}

	cg, err := sarama.NewConsumerGroup(cfg.Brokers, cfg.GroupID, cfg.SaramaConfig)
	if err != nil {
		return nil, err
	}

	// Значения по умолчанию
	if cfg.BatchSize <= 0 {
		cfg.BatchSize = 100
	}
	if cfg.BatchTimeout <= 0 {
		cfg.BatchTimeout = 5 * time.Second
	}

	prefix := fmt.Sprintf("consumer[%s] ", cfg.ConsumerID)
	logger := log.New(os.Stdout, prefix, log.LstdFlags|log.Lmsgprefix)

	c := &Consumer{
		topics:           cfg.Topics,
		consumerGroup:    cg,
		handler:          handler,
		deduplicator:     dedup,
		ready:            make(chan bool),
		id:               cfg.ConsumerID,
		batchSize:        cfg.BatchSize,
		batchTimeout:     cfg.BatchTimeout,
		flushOnRebalance: cfg.FlushOnRebalance,
		logger:           logger,
	}

	return c, nil
}

func (c *Consumer) Start(ctx context.Context, topics []string) error {
	// отдельная горутина для ошибок Сonsumer Group (полезно для диагностики)
	go func() {
		for err := range c.consumerGroup.Errors() {
			log.Printf("[consumer-group] error: %v", err)
		}
	}()

	c.wg.Add(1)
	go func() {
		defer c.wg.Done()
		for {
			if err := c.consumerGroup.Consume(ctx, topics, c); err != nil {
				c.logger.Printf("Consume error: %v", err)
			}
			if ctx.Err() != nil {
				return
			}
			c.ready = make(chan bool)
		}
	}()

	<-c.ready
	c.logger.Println("Consumer started with batching")
	return nil
}

// Stop gracefully останавливает
func (c *Consumer) Stop() error {
	close(c.ready)
	c.wg.Wait()
	c.logger.Println("Stopping...")
	return c.consumerGroup.Close()
}

func (c *Consumer) Setup(sarama.ConsumerGroupSession) error {
	close(c.ready)
	return nil
}

func (c *Consumer) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

func (c *Consumer) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	batch := make([]*sarama.ConsumerMessage, 0, c.batchSize)
	ticker := time.NewTicker(c.batchTimeout)
	defer ticker.Stop()

	flush := func() error {
		if len(batch) == 0 {
			return nil
		}

		// Дедупликация
		if c.deduplicator != nil {
			filtered := batch[:0]
			for _, msg := range batch {
				msgID := extractMessageID(msg)
				if msgID == "" || !c.deduplicator.Seen(msgID) {
					filtered = append(filtered, msg)
					if msgID != "" {
						c.deduplicator.MarkSeen(msgID)
					}
				} else {
					c.logger.Printf("Duplicate skipped: ID=%s, offset=%d", msgID, msg.Offset)
					session.MarkMessage(msg, "")
				}
			}
			batch = filtered
		}

		if len(batch) > 0 {
			if err := c.handler.HandleBatch(session.Context(), batch); err != nil {
				c.logger.Printf("Batch handle error: %v", err)
				return err
			}
		}

		for _, msg := range batch {
			session.MarkMessage(msg, "")
		}
		batch = batch[:0]
		return nil
	}

	for {
		select {
		case msg, ok := <-claim.Messages():
			if !ok {
				flush()
				return nil
			}
			batch = append(batch, msg)

			if len(batch) >= c.batchSize {
				if err := flush(); err != nil {
					return err
				}
				ticker.Reset(c.batchTimeout)
			}

		case <-ticker.C:
			if err := flush(); err != nil {
				return err
			}
			ticker.Reset(c.batchTimeout)

		case <-session.Context().Done():
			flush()
			return nil
		}
	}
}

// extractMessageID — из Key или Header
func extractMessageID(msg *sarama.ConsumerMessage) string {
	if len(msg.Key) > 0 {
		return string(msg.Key)
	}
	for _, h := range msg.Headers {
		if string(h.Key) == "message-id" {
			return string(h.Value)
		}
	}
	return ""
}
