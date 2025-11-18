package boot

import (
	"context"
	"errors"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/PechatnovVladimir/msa_big_tech/lib/closer"
	"github.com/PechatnovVladimir/msa_big_tech/lib/config"
	"github.com/PechatnovVladimir/msa_big_tech/lib/interceptors"
	"github.com/PechatnovVladimir/msa_big_tech/lib/kafka/consumer"
	"github.com/PechatnovVladimir/msa_big_tech/lib/postgres"
	"github.com/PechatnovVladimir/msa_big_tech/lib/postgres/transaction_manager"
	"github.com/PechatnovVladimir/msa_big_tech/lib/secrets"
	"google.golang.org/grpc"
	"log"
	"net"
	"strconv"
	"sync"
	"time"
)

type App struct {
	ctx             context.Context
	cfg             *config.Config
	secret          *secrets.Secrets
	lis             net.Listener
	grpcServer      *grpc.Server
	grpcRegister    func(*grpc.Server)
	db              *postgres.Connection
	tx              *transaction_manager.TransactionManager
	syncProducer    sarama.SyncProducer
	consumerManager *consumer.ConsumerManager
	Cl              *closer.Closer
}

type Option func(*App) error

func NewApp(ctx context.Context, opts ...Option) (*App, error) {
	app := &App{
		ctx: ctx,
	}
	for _, opt := range opts {
		if err := opt(app); err != nil {
			return nil, fmt.Errorf("failed to apply option: %w", err)
		}
	}

	if app.cfg == nil {
		return nil, errors.New("config is required")
	}

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(app.cfg.Grpc.Server.Port))

	if err != nil {
		return nil, fmt.Errorf("failed to listen: %w", err)
	}

	app.lis = lis

	app.Cl = &closer.Closer{}

	app.grpcServer = grpc.NewServer(
		interceptors.ServerInterceptors(app.cfg.Grpc.Server)...,
	)

	return app, nil
}

// WithConfigXXX конфиг с интегрированными секретами, ENV переменными, значениями по умолчанию
func WithConfigXXX(ctx context.Context, cfg Config) Option {
	return func(app *App) error {
		cl := config.NewConfigLoader().Configure(
			config.WithDefaults(cfg.DefaultValue),
			config.WithEnvBindings(cfg.EnvBinding),
		)

		if cfg.SecretProvider != nil {
			secret := secrets.NewSecrets(cfg.SecretProvider)
			if secret.IsSet() {
				cl.WithSecretsProvider(secret)
				cl.AddSecretKeys(cfg.SecretKeys...)
			}
		}

		cfg, err := cl.LoadConfig(cfg.ConfigFile)
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}

		app.cfg = cfg
		return nil
	}
}

func (app *App) Run(ctx context.Context) error {
	if app.grpcRegister != nil {
		app.grpcRegister(app.grpcServer)
	}

	if app.consumerManager != nil {
		err := app.consumerManager.StartAll(ctx)
		if err != nil {
			return err
		}
	}

	app.Cl.Add(func(ctx context.Context) error {
		log.Println("grpc server stopped")
		app.grpcServer.GracefulStop()
		return nil
	})

	var wg sync.WaitGroup
	wg.Add(1)

	// Run gRPC server in a goroutine
	go func() {
		defer wg.Done()
		log.Printf("%s %s - gRPC server listening on :%s", app.cfg.App.Name, app.cfg.App.Version, strconv.Itoa(app.cfg.Grpc.Server.Port))
		if err := app.grpcServer.Serve(app.lis); err != nil && err != grpc.ErrServerStopped {
			log.Fatalf("gRPC server failed: %v", err)
		}
	}()

	//ждем сигнал на останов
	<-ctx.Done()

	log.Println("server: shutting down server gracefully")

	// Create a context with a 20-second timeout
	shutdownCtx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	err := app.Cl.CloseAll(shutdownCtx)
	if err != nil {
		log.Printf("failed to close clients: %v", err)
		return fmt.Errorf("closer: %v", err)
	}

	log.Println("server: shutdown completed")

	wg.Wait()
	return nil
}
