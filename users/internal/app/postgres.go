package app

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	PostgresHost     = "localhost"
	PostgresPort     = "5432"
	PostgresDB       = "postgres-users"
	PostgresUser     = "postgres-users-user"
	PostgresPassword = "postgres-users-psw"
)

func DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		PostgresUser, PostgresPassword, PostgresHost, PostgresPort, PostgresDB,
	)
}
func NewPostgresConnection(ctx context.Context) (*pgxpool.Pool, error) {
	cfg, err := pgxpool.ParseConfig(DSN())
	if err != nil {
		return nil, fmt.Errorf("parse config: %w", err)
	}

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, fmt.Errorf("pgxpool connect: %w", err)
	}

	return pool, nil
}
