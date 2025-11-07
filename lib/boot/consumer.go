package boot

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/lib/kafka/consumer"
	"log"
	"time"
)

func (app *App) ConsumerManager(ctx context.Context) *consumer.ConsumerManager {
	manager := consumer.NewManager()
	app.consumerManager = manager
	app.Cl.Add(func(ctx context.Context) error {
		log.Println("consumer stopped")
		app.consumerManager.StopAll()
		return nil
	})

	return app.consumerManager
}

func (app *App) Consumer(ctx context.Context, topics []string, dedup consumer.Deduplicator, h consumer.ConsumerHandler, consumerID string, groupID string) (*consumer.Consumer, error) {
	cfg := consumer.Config{
		Brokers:      []string{app.cfg.KafkaConsumer.Brokers},
		Topics:       topics,
		GroupID:      groupID,
		BatchSize:    10,
		BatchTimeout: 5 * time.Second,
		ConsumerID:   consumerID,
	}

	c, err := consumer.NewConsumer(cfg, h, dedup)
	if err != nil {
		return nil, err
	}

	return c, nil
}
