package app

import (
	"context"
	seh "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/adapters/chateventshandler"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/adapters/userprovider"
	serv "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/controllers/chat/grpc/v1"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/modules/outbox"
	chatRepo "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/repositories/chat"
	outboxRepo "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/repositories/outbox"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat"
	pb "github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
	"github.com/PechatnovVladimir/msa_big_tech/lib/boot"
	"google.golang.org/grpc"
	"log"
	"os"
	"time"
)

func Start(ctx context.Context) (err error) {

	ctx = context.WithValue(ctx, "CurrentUser", os.Getenv("CurrentUser"))

	app, err := boot.NewApp(ctx,
		boot.WithConfig(ctx, "./config/config.yaml"),
	)

	if err != nil {
		return err
	}

	conn, txManager, err := app.Postgres(ctx)

	if err != nil {
		return err
	}
	defer conn.Close()

	producer, err := app.SyncProducer(ctx)
	if err != nil {
		log.Fatal(err)
	}

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

	service := serv.New(serv.Deps{
		ChatUseCase: chatUseCase,
	})

	app.RegisterGRPC(func(srv *grpc.Server) {
		pb.RegisterChatServiceServer(srv, service)
	})

	err = app.Run(ctx)
	if err != nil {
		return err
	}

	return nil

}
