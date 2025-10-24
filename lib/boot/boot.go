package boot

import (
	"context"
	"errors"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/PechatnovVladimir/msa_big_tech/lib/config"
	"github.com/PechatnovVladimir/msa_big_tech/lib/interceptors"
	"github.com/PechatnovVladimir/msa_big_tech/lib/postgres"
	"github.com/PechatnovVladimir/msa_big_tech/lib/postgres/transaction_manager"
	"github.com/PechatnovVladimir/msa_big_tech/lib/secrets"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"os"
	"os/signal"
	"strconv"
	"sync"
	"syscall"
)

type App struct {
	ctx          context.Context
	cfg          *config.Config
	secret       *secrets.Secrets
	grpcServer   *grpc.Server
	grpcRegister func(*grpc.Server)
	db           *postgres.Connection
	tx           *transaction_manager.TransactionManager
	syncProducer sarama.SyncProducer
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

	app.grpcServer = grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.ChainUnaryInterceptor(interceptors.ProtoValidate),
	)

	return app, nil
}

func WithConfig(ctx context.Context, cfgFile string) Option {
	return func(app *App) error {
		cfg, err := config.LoadConfig(ctx, cfgFile)
		if err != nil {
			return fmt.Errorf("failed to load config: %w", err)
		}
		app.cfg = cfg
		return nil
	}
}

func WithSecret(secret *secrets.Secrets) Option {
	return func(app *App) error {
		app.secret = secret
		return nil
	}
}

func (app *App) Run(ctx context.Context) error {
	if app.grpcRegister != nil {
		app.grpcRegister(app.grpcServer)
	}

	lis, err := net.Listen("tcp", ":"+strconv.Itoa(app.cfg.Grpc.Port))
	if err != nil {
		return fmt.Errorf("failed to listen: %w", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)

	// Run gRPC server in a goroutine
	go func() {
		defer wg.Done()
		log.Printf("%s %s - gRPC server listening on :%s", app.cfg.App.Name, app.cfg.App.Version, strconv.Itoa(app.cfg.Grpc.Port))
		if err := app.grpcServer.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			log.Fatalf("gRPC server failed: %v", err)
		}
	}()

	// Handle graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	select {
	case <-ctx.Done():
		log.Println("Context cancelled, shutting down")
	case sig := <-sigChan:
		log.Printf("Received signal %v, shutting down", sig)
	}

	app.grpcServer.GracefulStop()

	if app.db != nil {
		app.db.Close()
	}
	wg.Wait()
	return nil
}
