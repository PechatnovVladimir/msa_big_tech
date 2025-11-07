package boot

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/PechatnovVladimir/msa_big_tech/lib/kafka"
	"log"
	"strings"
)

func (app *App) SyncProducer(ctx context.Context) (sarama.SyncProducer, error) {
	if app.syncProducer == nil {
		producer, err := kafka.NewSyncProducer(strings.Split(app.cfg.KafkaProducer.Brokers, ","), nil)
		if err != nil {
			return nil, err
		}
		app.syncProducer = producer
	}

	app.Cl.Add(func(ctx context.Context) error {
		log.Println("sync producer closed")
		app.syncProducer.Close()
		return nil
	})

	return app.syncProducer, nil
}
