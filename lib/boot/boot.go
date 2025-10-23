package boot

import (
	"context"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/lib/config"
	"github.com/PechatnovVladimir/msa_big_tech/lib/postgres"
	"github.com/PechatnovVladimir/msa_big_tech/lib/secrets"
	"google.golang.org/grpc"
	"time"
)

type App struct {
	ctx          context.Context
	cfg          config.Config
	secret       secrets.Secrets
	grpcServer   *grpc.Server
	grpcRegister func(*grpc.Server)
	db           *postgres.Connection
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
	app.grpcServer = grpc.NewServer()

	return app, nil
}

func WithConfig(cfg config.Config) Option {
	return func(app *App) error {
		app.cfg = cfg
		return nil
	}
}

func (app *App) Postgress(ctx context.Context) (*postgres.Connection, error) {
	if app.db == nil {
		conn, err := postgres.NewConnectionPool(ctx, app.cfg.Postgres.DSN(),
			postgres.WithMaxConnIdleTime(time.Minute),
			postgres.WithMinConnectionsCount(3),
			postgres.WithMaxConnectionsCount(10),
		)
		if err != nil {
			return nil, err
		}
		app.db = conn
	}
	return app.db, nil
}
