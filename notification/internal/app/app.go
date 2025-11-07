package app

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/lib/boot"
	"github.com/PechatnovVladimir/msa_big_tech/lib/kafka/consumer"
	inboxRepo "github.com/PechatnovVladimir/msa_big_tech/notification/internal/app/repository/inbox"
	"github.com/PechatnovVladimir/msa_big_tech/notification/internal/app/usecase/inbox"
	"log"
	"os/signal"
	"syscall"
	"time"
)

var (
	Topics        = []string{"chat.message.sent"}
	ConsumerGroup = "notification-inbox-consumer-group"
	ConsumerID    = "notification-service" // Уникальный для каждого инстанса нашего приложения
)

var (
	defaultConfigValue = map[string]interface{}{
		"app.mode":          "dev",
		"app.name":          "NOTIFICATION-SERVICE",
		"app.version":       "0.0.1",
		"postgres.host":     "localhost",
		"postgres.port":     "5435",
		"postgres.user":     "postgres-notification-user",
		"postgres.password": "postgres-notification-psw",
		"postgres.database": "postgres-notification",
		"postgres.sslmode":  "disable",
		//"grpc.port":              "50052",
		//"grpc.host":              "localhost",
		//"kafka_producer.brokers": "localhost:9092",
		"kafka_consumer.brokers": "localhost:9092",
	}

	envConfigBinding = map[string]string{
		"app.mode":          "NOTIFICATION_APP_MODE",
		"app.name":          "NOTIFICATION_APP_NAME",
		"app.version":       "NOTIFICATION_APP_VERSION",
		"postgres.host":     "NOTIFICATION_POSTGRES_HOST",
		"postgres.port":     "NOTIFICATION_POSTGRES_PORT",
		"postgres.user":     "NOTIFICATION_POSTGRES_USER",
		"postgres.password": "NOTIFICATION_POSTGRES_PASSWORD",
		"postgres.database": "NOTIFICATION_POSTGRES_DATABASE",
		"postgres.sslmode":  "NOTIFICATION_POSTGRES_SSLMODE",
		//"grpc.port":              "NOTIFICATION_GRPC_PORT",
		//"grpc.host":              "NOTIFICATION_GRPC_HOST",
		//"kafka_producer.brokers": "NOTIFICATION_PRODUCER_BROKERS",
		"kafka_consumer.brokers": "NOTIFICATION_CONSUMER_BROKERS",
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

	//конфигурируем приложение
	app, err := boot.NewApp(ctx,
		boot.WithConfigXXX(ctx, config),
	)

	if err != nil {
		return err
	}
	defer app.Cl.CloseAll(context.TODO())

	//коннект к БД и менеджер транзакций
	conn, txManager, err := app.Postgres(ctx)

	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//репозиторий inbox
	inboxRepo := inboxRepo.NewRepository(txManager)

	//воркер по разбору inbox
	notificator := inbox.NewNotificator()
	worker := inbox.NewInboxMessageWorker(inboxRepo, txManager, notificator,
		inbox.WithBatchSize(10),
	)
	go worker.Run(ctx)

	manager := app.ConsumerManager(ctx)

	handler := inbox.New(inbox.Deps{InboxRepo: inboxRepo})
	deduplicator := consumer.NewInMemory(ctx, 24*time.Hour)

	c1, err := app.Consumer(ctx, Topics, deduplicator, handler, ConsumerID, ConsumerGroup)
	if err != nil {
		log.Fatal(err)
	}
	manager.Add(c1)

	handler2 := inbox.New(inbox.Deps{InboxRepo: inboxRepo})
	deduplicator2 := consumer.NewInMemory(ctx, 24*time.Hour)

	c2, err := app.Consumer(ctx, []string{"social.friend.request", "social.friend.updated"}, deduplicator2, handler2, "notification-service-social", ConsumerGroup)
	if err != nil {
		log.Fatal(err)
	}

	manager.Add(c2)

	err = app.Run(ctx)
	if err != nil {
		return err
	}

	return nil
}
