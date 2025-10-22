package app

import (
	"context"
	"fmt"
	seh "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/adapters/chateventshandler"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/adapters/userprovider"
	chatGPRS "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/controllers/chat/grpc"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/controllers/chat/grpc/v1"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/modules/outbox"
	chatRepo "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/repositories/chat"
	outboxRepo "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/repositories/outbox"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat"
	"github.com/PechatnovVladimir/msa_big_tech/lib/config"
	"github.com/PechatnovVladimir/msa_big_tech/pkg/kafka"
	connection "github.com/PechatnovVladimir/msa_big_tech/pkg/postgres"
	"github.com/PechatnovVladimir/msa_big_tech/pkg/postgres/transaction_manager"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

func Run(ctx context.Context, cfg *config.Config) (err error) {
	if cfg == nil {
		log.Fatal("config is nil")
	}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ctx = context.WithValue(ctx, "CurrentUser", os.Getenv("CurrentUser"))

	//соединение
	conn, err := connection.NewConnectionPool(ctx, cfg.Postgres.DSN(),
		connection.WithMaxConnIdleTime(time.Minute),
		connection.WithMinConnectionsCount(3),
		connection.WithMaxConnectionsCount(10),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	producer, err := kafka.NewSyncProducer(strings.Split(cfg.KafkaProducer.Brokers, ","), nil)
	if err != nil {
		log.Fatal(err)
	}

	//менеджер транзакций
	txManager := transaction_manager.New(conn)

	userServiceProvider := userprovider.New()
	chatRepository := chatRepo.NewRepository(txManager)
	outboxRepository := outboxRepo.NewRepository(txManager)

	chatEventsHandler := seh.NewKafkaBatchHandler(producer,
		seh.WithMaxBatchSize(100),
		seh.WithTopic(),
	)

	worker := outbox.NewOutboxChatWorker(outboxRepository, txManager, chatEventsHandler,
		outbox.WithBatchSize(10),
		outbox.WithMaxRetry(10),
		outbox.WithRetryInterval(30*time.Second),
		outbox.WithWindow(time.Hour),
	)

	go worker.Run(ctx)

	outboxProcessor := outbox.NewProcessor(outbox.Deps{Repository: outboxRepository})

	chatUseCase := chat.New(chat.Deps{
		UserProvider:       userServiceProvider,
		ChatRepo:           chatRepository,
		TransactionManager: txManager,
		OutboxRepo:         outboxProcessor,
	})

	//grpc
	grpcServer, err := chatGPRS.New(v1.Deps{
		ChatUseCase: chatUseCase,
		Cfg:         &cfg.Grpc,
	})

	if err != nil {
		return fmt.Errorf("%s - grpc.New: %w", cfg.App.Name, err)
	}

	log.Println(fmt.Sprintf("%s started on port %d", cfg.App.Name, cfg.Grpc.Port))

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig // wait signal

	grpcServer.Close()

	log.Println(fmt.Sprintf("%s stopped!", cfg.App.Name))

	return nil

}
