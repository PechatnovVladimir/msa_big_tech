package app

import (
	"context"
	"fmt"
	connection "github.com/PechatnovVladimir/msa_big_tech/pkg/postgres"
	tx "github.com/PechatnovVladimir/msa_big_tech/pkg/postgres/transaction_manager"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/adapters/authprovider"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/adapters/friend_request_events_handler"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/adapters/userprovider"
	socialGPRS "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/controllers/social/grpc"
	v1 "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/controllers/social/grpc/v1"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/modules/outbox"
	outboxRepo "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/repositories/outbox"
	socialRepo "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/repositories/social"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(ctx context.Context) (err error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
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
	txManager := tx.New(conn)

	authServiceAdapter := authprovider.New()
	userServiceAdapter := userprovider.New()
	socialRepository := socialRepo.New(txManager)
	outboxRepository := outboxRepo.NewRepository(txManager)

	friendRequestEventsHandler := friend_request_events_handler.NewKafkaFriendRequestBatchHandler()

	worker := outbox.NewOutboxFriendRequestWorker(outboxRepository, txManager, friendRequestEventsHandler,
		outbox.WithBatchSize(10),
		outbox.WithMaxRetry(10),
		outbox.WithRetryInterval(30*time.Second),
		outbox.WithWindow(time.Hour),
	)

	go worker.Run(ctx)

	outboxProcessor := outbox.NewProcessor(outbox.Deps{Repository: outboxRepository})

	socialUseCase := social.New(social.Deps{
		AuthProvider:       authServiceAdapter,
		UserProvider:       userServiceAdapter,
		SocialRepo:         socialRepository,
		TransactionManager: txManager,
		OutboxRepo:         outboxProcessor,
	})

	//grpc
	grpcServer, err := socialGPRS.New(v1.Deps{
		SocialUseCase: socialUseCase,
	})
	if err != nil {
		return fmt.Errorf("authGRPC.New: %w", err)
	}

	log.Println("Social service started!")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig // wait signal

	grpcServer.Close()

	log.Println("Social service stopped!!!!")

	return nil

}
