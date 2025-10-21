package app

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/notification/internal/app/controllers/kafkaconsumer"
	inboxRepo "github.com/PechatnovVladimir/msa_big_tech/notification/internal/app/repository/inbox"
	"github.com/PechatnovVladimir/msa_big_tech/notification/internal/app/usecase/inbox"
	connection "github.com/PechatnovVladimir/msa_big_tech/pkg/postgres"
	"github.com/PechatnovVladimir/msa_big_tech/pkg/postgres/transaction_manager"
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
	KafkaConsumerName  = "notification-service-2" // Уникальный для каждого инстанса нашего приложения
)

func Run(ctx context.Context) (err error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	dedup := kafkaconsumer.NewInMemoryDeduper(ctx, 24*time.Hour)

	//соединение
	conn, err := connection.NewConnectionPool(ctx, DSN(),
		connection.WithMaxConnIdleTime(time.Minute),
		connection.WithMinConnectionsCount(3),
		connection.WithMaxConnectionsCount(10),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//менеджер транзакций
	txManager := transaction_manager.New(conn)

	inboxRepo := inboxRepo.NewRepository(txManager)

	usecase := inbox.New(inbox.Deps{InboxRepo: inboxRepo})

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

	notificator := inbox.NewNotificator()

	worker := inbox.NewInboxMessageWorker(inboxRepo, txManager, notificator,
		inbox.WithBatchSize(10),
		inbox.WithMaxRetry(10),
		inbox.WithRetryInterval(10*time.Second),
		inbox.WithWindow(time.Hour),
	)

	go worker.Run(ctx)

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
