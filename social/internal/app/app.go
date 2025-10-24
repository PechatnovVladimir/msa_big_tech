package app

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/lib/boot"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/adapters/authprovider"
	seh "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/adapters/socialeventshandler"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/adapters/userprovider"
	serv "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/controllers/social/grpc/v1"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/modules/outbox"
	outboxRepo "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/repositories/outbox"
	socialRepo "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/repositories/social"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social"
	pb "github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
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

	authServiceAdapter := authprovider.New()
	userServiceAdapter := userprovider.New()
	socialRepository := socialRepo.New(txManager)
	outboxRepository := outboxRepo.NewRepository(txManager)

	socialEventsHandler := seh.NewKafkaBatchHandler(producer,
		seh.WithMaxBatchSize(10),
		seh.WithTopic(),
	)

	worker := outbox.NewOutboxFriendRequestWorker(outboxRepository, txManager, socialEventsHandler,
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

	service := serv.New(serv.Deps{
		SocialUseCase: socialUseCase,
	})

	app.RegisterGRPC(func(srv *grpc.Server) {
		pb.RegisterSocialServiceServer(srv, service)
	})

	err = app.Run(ctx)
	if err != nil {
		return err
	}

	return nil

}
