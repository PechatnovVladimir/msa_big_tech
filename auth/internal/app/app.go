package app

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/adapters/userservice"
	serv "github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/controllers/auth/grpc/v1"
	pb "github.com/PechatnovVladimir/msa_big_tech/auth/pkg/proto/api/auth/v1"
	"github.com/PechatnovVladimir/msa_big_tech/lib/boot"
	"google.golang.org/grpc"
	//authRepo "github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/repositories/auth/inmemory"
	authRepo "github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/repositories/auth"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/usecases/auth"
)

var (
	defaultConfigValue = map[string]interface{}{
		"app.mode":               "dev",
		"app.name":               "AUTH-SERVICE",
		"app.version":            "0.0.1",
		"postgres.host":          "localhost",
		"postgres.port":          "5431",
		"postgres.user":          "postgres-auth-user",
		"postgres.password":      "postgres-auth-psw",
		"postgres.database":      "postgres-auth-user",
		"postgres.sslmode":       "disable",
		"grpc.port":              "50051",
		"grpc.host":              "localhost",
		"kafka_producer.brokers": "localhost:9092",
		"kafka_consumer.brokers": "localhost:9092",
	}

	envConfigBinding = map[string]string{
		"app.mode":               "AUTH_APP_MODE",
		"app.name":               "AUTH_APP_NAME",
		"app.version":            "AUTH_APP_VERSION",
		"postgres.host":          "AUTH_POSTGRES_HOST",
		"postgres.port":          "AUTH_POSTGRES_PORT",
		"postgres.user":          "AUTH_POSTGRES_USER",
		"postgres.password":      "AUTH_POSTGRES_PASSWORD",
		"postgres.database":      "AUTH_POSTGRES_DATABASE",
		"postgres.sslmode":       "AUTH_POSTGRES_SSLMODE",
		"grpc.port":              "AUTH_GRPC_PORT",
		"grpc.host":              "AUTH_GRPC_HOST",
		"kafka_producer.brokers": "AUTH_PRODUCER_BROKERS",
		"kafka_consumer.brokers": "AUTH_CONSUMER_BROKERS",
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
	app, err := boot.NewApp(ctx,
		boot.WithConfigXXX(ctx, config),
	)

	if err != nil {
		return err
	}

	userServiceAdapter := userservice.New()

	authRepository := authRepo.New()

	authUseCase := auth.New(auth.Deps{
		UserService: userServiceAdapter,
		AuthRepo:    authRepository,
	})

	service := serv.New(serv.Deps{
		AuthUseCase: authUseCase,
	})

	app.RegisterGRPC(func(srv *grpc.Server) {
		pb.RegisterAuthServiceServer(srv, service)
	})

	err = app.Run(ctx)
	if err != nil {
		return err
	}

	return nil

}
