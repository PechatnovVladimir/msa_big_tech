package app

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/notification/internal/app/controllers/kafkaconsumer"
	"github.com/PechatnovVladimir/msa_big_tech/notification/internal/app/usecase/inbox"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	KafkaBrokers       = "localhost:9092"
	KafkaTopicName     = "chat.message.sent"
	KafkaConsumerGroup = "notification-inbox-consumer-group"
	KafkaConsumerName  = "notification-service-1" // Уникальный для каждого инстанса нашего приложения
)

func Run(ctx context.Context) (err error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// adapters/repository
	dedup := kafkaconsumer.NewInMemoryDeduper(ctx, 24*time.Hour)

	usecase := inbox.New(inbox.Deps{InboxRepo: nil})

	consumer, err := kafkaconsumer.NewInboxConsumer([]string{KafkaBrokers},
		KafkaConsumerGroup,
		KafkaConsumerName,
		dedup,
		usecase,
	)
	if err != nil {
		log.Fatal(err)
	}

	defer consumer.Close()

	if err := consumer.Run(ctx, KafkaTopicName); err != nil && ctx.Err() == nil {
		log.Println("consumer stopped", err)
	}

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig // wait signal

	consumer.Close()

	log.Println("Notification service stopped!!!!")

	return nil
}
