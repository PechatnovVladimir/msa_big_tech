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
	"os/signal"
	"syscall"
	"time"
)

var (
	defaultConfigValue = map[string]interface{}{
		"app.mode":               "dev",
		"app.name":               "SOCIAL-SERVICE",
		"app.version":            "0.0.1",
		"postgres.host":          "localhost",
		"postgres.port":          "5433",
		"postgres.user":          "postgres-soc-user",
		"postgres.password":      "postgres-soc-psw",
		"postgres.database":      "postgres-soc",
		"postgres.sslmode":       "disable",
		"grpc.server.port":       "50053",
		"grpc.server.host":       "localhost",
		"kafka_producer.brokers": "localhost:9092",
		"kafka_consumer.brokers": "localhost:9092",
	}

	envConfigBinding = map[string]string{
		"app.mode":               "SOC_APP_MODE",
		"app.name":               "SOC_APP_NAME",
		"app.version":            "SOC_APP_VERSION",
		"postgres.host":          "SOC_POSTGRES_HOST",
		"postgres.port":          "SOC_POSTGRES_PORT",
		"postgres.user":          "SOC_POSTGRES_USER",
		"postgres.password":      "SOC_POSTGRES_PASSWORD",
		"postgres.database":      "SOC_POSTGRES_DATABASE",
		"postgres.sslmode":       "SOC_POSTGRES_SSLMODE",
		"grpc.server.port":       "SOC_GRPC_PORT",
		"grpc.server.host":       "SOC_GRPC_HOST",
		"kafka_producer.brokers": "SOC_PRODUCER_BROKERS",
		"kafka_consumer.brokers": "SOC_CONSUMER_BROKERS",
	}
)

var (
	config = boot.Config{
		ConfigFile:   "./config/config.yaml",
		DefaultValue: defaultConfigValue,
		EnvBinding:   envConfigBinding,
		//SecretProvider: secrets.NewEnvProvider()
		//SecretProvider: secrets.NewVaultProvider("http://localhost:8200", "secret/data/users-service"),
		//SecretProvider: secrets.NewFileProvider("/Users/pvv/Educ/Balun/msa_big_tech/secrets_example.yaml"),
		//SecretProvider: secrets.NewCompositeProvider("http://localhost:8200", "secret/data/users-service", "/Users/pvv/Educ/Balun/msa_big_tech/secrets_example.yaml"),
		//SecretKeys:     []string{"postgres.user", "postgres.password"},
	}
)

func Start(ctx context.Context) (err error) {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	//конфигурируем приложение
	app, err := boot.NewApp(ctx,
		boot.WithConfigXXX(ctx, config),
	)

	if err != nil {
		return err
	}
	defer app.Cl.CloseAll(context.TODO())

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
