package app

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/lib/boot"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/adapters/userprovider"
	serv "github.com/PechatnovVladimir/msa_big_tech/users/internal/app/controllers/users/grpc/v1"
	repo "github.com/PechatnovVladimir/msa_big_tech/users/internal/app/repositories/users"
	uc "github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users"
	pb "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
	"google.golang.org/grpc"
)

var (
	defaultConfigValue = map[string]interface{}{
		"app.mode":               "dev",
		"app.name":               "USERS-SERVICE",
		"app.version":            "0.0.1",
		"postgres.host":          "localhost",
		"postgres.port":          "5434",
		"postgres.user":          "postgres-users-user",
		"postgres.password":      "postgres-users-psw",
		"postgres.database":      "postgres-users",
		"postgres.sslmode":       "disable",
		"grpc.server.port":       "50054",
		"grpc.server.host":       "localhost",
		"kafka_producer.brokers": "localhost:9092",
		"kafka_consumer.brokers": "localhost:9092",
	}

	envConfigBinding = map[string]string{
		"app.mode":               "USERS_APP_MODE",
		"app.name":               "USERS_APP_NAME",
		"app.version":            "USERS_APP_VERSION",
		"postgres.host":          "USERS_POSTGRES_HOST",
		"postgres.port":          "USERS_POSTGRES_PORT",
		"postgres.user":          "USERS_POSTGRES_USER",
		"postgres.password":      "USERS_POSTGRES_PASSWORD",
		"postgres.database":      "USERS_POSTGRES_DATABASE",
		"postgres.sslmode":       "USERS_POSTGRES_SSLMODE",
		"grpc.port":              "USERS_GRPC_PORT",
		"grpc.host":              "USERS_GRPC_HOST",
		"kafka_producer.brokers": "USERS_PRODUCER_BROKERS",
		"kafka_consumer.brokers": "USERS_CONSUMER_BROKERS",
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
	app, err := boot.NewApp(ctx,
		boot.WithConfigXXX(ctx, config),
	)

	if err != nil {
		return err
	}

	conn, txManager, err := app.Postgres(ctx)

	if err != nil {
		return err
	}
	defer conn.Close()

	userProvider := userprovider.New()

	repository := repo.New(txManager)

	userUseCase := uc.New(uc.Deps{
		UserRepo:     repository,
		UserProvider: userProvider,
	})

	service := serv.New(userUseCase)

	app.RegisterGRPC(func(srv *grpc.Server) {
		pb.RegisterUserServiceServer(srv, service)
	})

	err = app.Run(ctx)
	if err != nil {
		return err
	}

	return nil
}
