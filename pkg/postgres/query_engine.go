package postgres

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

// QueryEngineProvider - smths that gives us QueryEngine
type QueryEngineProvider interface {
	GetQueryEngine(ctx context.Context) QueryEngine
}

// PgxCommonAPI - pgx common api
type PgxCommonAPI interface {
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
}

// PgxCommonScanAPI улучшенный PgxCommonAPI
type PgxCommonScanAPI interface {
	// Getx - aka QueryRow
	Getx(ctx context.Context, dest interface{}, sqlizer Sqlizer) error
	// Selectx - aka Query
	Selectx(ctx context.Context, dest interface{}, sqlizer Sqlizer) error
	// Execx - aka Exec
	Execx(ctx context.Context, sqlizer Sqlizer) (pgconn.CommandTag, error)
}

// PgxExtendedAPI - ...
type PgxExtendedAPI interface {
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	CopyFrom(ctx context.Context, tableName pgx.Identifier, columnNames []string, rowSrc pgx.CopyFromSource) (int64, error)
}

// QueryEngine is a common database query interface.
type QueryEngine interface {
	PgxCommonAPI
	PgxCommonScanAPI
	PgxExtendedAPI
}
