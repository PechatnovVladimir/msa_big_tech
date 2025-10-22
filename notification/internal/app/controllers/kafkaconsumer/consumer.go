package kafkaconsumer

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/PechatnovVladimir/msa_big_tech/notification/internal/app/usecase/inbox"
	"log"
	"regexp"
	"strings"
	"time"
)

const headerEventID = "event_id"

var topicRE = regexp.MustCompile(`^[A-Za-z0-9._-]{1,249}$`)

func extractID(msg *sarama.ConsumerMessage) (string, bool) {
	for _, h := range msg.Headers {
		if strings.EqualFold(string(h.Key), headerEventID) && len(h.Value) > 0 {
			return string(h.Value), true
		}
	}
	return "", false
}

type InboxConsumer struct {
	group        sarama.ConsumerGroup
	dedup        Deduper
	handler      *inbox.Service
	batchSize    int
	batchTimeout time.Duration
	consumerName string
}

func NewInboxConsumer(brokers []string, groupID string, consumerName string, dedup Deduper, h *inbox.Service) (*InboxConsumer, error) {
	cfg := sarama.NewConfig()
	cfg.Version = sarama.V2_5_0_0
	cfg.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRange()
	cfg.Consumer.Offsets.Initial = sarama.OffsetOldest
	cfg.Consumer.Return.Errors = true // ОБЯЗАТЕЛЬНО ЧИТАЕМ cg.group.Errors()

	cg, err := sarama.NewConsumerGroup(brokers, groupID, cfg)
	if err != nil {
		return nil, err
	}

	return &InboxConsumer{
		group:        cg,
		dedup:        dedup,
		handler:      h,
		batchSize:    128,
		batchTimeout: 300 * time.Millisecond,
		consumerName: consumerName,
	}, nil
}

func (c *InboxConsumer) Close() error { return c.group.Close() }

func (c *InboxConsumer) Run(ctx context.Context, topics ...string) error {
	for _, t := range topics {
		if !topicRE.MatchString(t) {
			return sarama.ConfigurationError("invalid topic: " + t)
		}
	}

	// отдельная горутина для ошибок Сonsumer Group (полезно для диагностики)
	go func() {
		for err := range c.group.Errors() {
			log.Printf("[consumer-group] error: %v", err)
		}
	}()

	handler := &consumerGroupHandler{c: c}
	for {
		if err := c.group.Consume(ctx, topics, handler); err != nil {
			return err
		}
		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
}

type consumerGroupHandler struct{ c *InboxConsumer }

func (h *consumerGroupHandler) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (h *consumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }

// ConsumeClaim вызывается отдельно на КАЖДУЮ партицию (важно для порядка сообщений)
func (h *consumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	batch := make([]*sarama.ConsumerMessage, 0, h.c.batchSize)

	timer := time.NewTimer(h.c.batchTimeout)
	defer timer.Stop()

	flush := func() {
		if len(batch) == 0 {
			return
		}
		for _, msg := range batch {
			id, ok := extractID(msg)
			if !ok {
				// без ID не можем гарантировать идемпотентность — безопаснее скипнуть и скоммитить,
				// либо отправить в DLQ.
				log.Printf("skip message without ID (commit offset): topic=%s p=%d off=%d",
					msg.Topic, msg.Partition, msg.Offset)
				sess.MarkMessage(msg, "")
				continue
			}

			// дедупликация: если уже обработано — просто коммитим offset и идём дальше
			if h.c.dedup.Seen(id) {
				sess.MarkMessage(msg, "")
				continue
			}

			// бизнес-обработка (идемпотентная)
			if err := h.c.handler.SaveMessage(sess.Context(), msg); err != nil {
				// обработка упала: НЕ коммитим offset -> Kafka переотправит (at-least-once)
				log.Printf("handle failed (will retry): id=%s topic=%s p=%d off=%d err=%v",
					id, msg.Topic, msg.Partition, msg.Offset, err)
				continue
			}

			// успех: пометили как Done в дедупе и коммитим offset
			h.c.dedup.MarkDone(id)
			sess.MarkMessage(msg, "")
		}
		batch = batch[:0]
	}

	for {
		select {
		case <-sess.Context().Done():
			return nil

		case m, ok := <-claim.Messages():
			if !ok {
				flush()
				return nil
			}
			batch = append(batch, m)
			if len(batch) >= h.c.batchSize {
				flush()
				if !timer.Stop() {
					select {
					case <-timer.C:
					default:
					}
				}
				timer.Reset(h.c.batchTimeout)
			}

		case <-timer.C:
			flush()
			timer.Reset(h.c.batchTimeout)
		}
	}
}
