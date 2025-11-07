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
	"os/signal"
	"syscall"
	"time"
)

var (
	defaultConfigValue = map[string]interface{}{
		"app.mode":               "dev",
		"app.name":               "CHAT-SERVICE",
		"app.version":            "0.0.1",
		"postgres.host":          "localhost",
		"postgres.port":          "5432",
		"postgres.user":          "postgres-chat-user",
		"postgres.password":      "postgres-chat-psw",
		"postgres.database":      "postgres-chat",
		"postgres.sslmode":       "disable",
		"grpc.server.port":       "50052",
		"grpc.server.host":       "localhost",
		"kafka_producer.brokers": "localhost:9092",
		"kafka_consumer.brokers": "localhost:9092",
	}

	envConfigBinding = map[string]string{
		"app.mode":               "CHAT_APP_MODE",
		"app.name":               "CHAT_APP_NAME",
		"app.version":            "CHAT_APP_VERSION",
		"postgres.host":          "CHAT_POSTGRES_HOST",
		"postgres.port":          "CHAT_POSTGRES_PORT",
		"postgres.user":          "CHAT_POSTGRES_USER",
		"postgres.password":      "CHAT_POSTGRES_PASSWORD",
		"postgres.database":      "CHAT_POSTGRES_DATABASE",
		"postgres.sslmode":       "CHAT_POSTGRES_SSLMODE",
		"grpc.server.port":       "CHAT_GRPC_PORT",
		"grpc.server.host":       "CHAT_GRPC_HOST",
		"kafka_producer.brokers": "CHAT_PRODUCER_BROKERS",
		"kafka_consumer.brokers": "CHAT_CONSUMER_BROKERS",
	}
)

var (
	config = boot.Config{
		ConfigFile:   "./config/config.yaml",
		DefaultValue: defaultConfigValue,
		EnvBinding:   envConfigBinding,
		//SecretProvider: secrets.NewEnvProvider(),
		//SecretProvider: secrets.NewVaultProvider("http://localhost:8200", "secret/data/users-service"),
		//SecretProvider: secrets.NewFileProvider("/Users/pvv/Educ/Balun/msa_big_tech/secrets_example.yaml"),
		//SecretProvider: secrets.NewCompositeProvider("http://localhost:8200", "secret/data/users-service", "/Users/pvv/Educ/Balun/msa_big_tech/secrets_example.yaml"),
		//SecretKeys: []string{"postgres.user", "postgres.password"},
	}
)

func Start(ctx context.Context) (err error) {
	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

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

	app.Cl.Add(func(ctx context.Context) error {
		log.Println("worker outbox chat closed")
		//TODO тут остановку worker надо
		return nil
	})

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
