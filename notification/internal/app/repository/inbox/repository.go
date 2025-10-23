package inbox

import (
	"github.com/Masterminds/squirrel"
	"github.com/PechatnovVladimir/msa_big_tech/lib/postgres"
)

type Repository struct {
	db postgres.QueryEngineProvider
	sb squirrel.StatementBuilderType
}

// NewRepository конструктор Repository
func NewRepository(p postgres.QueryEngineProvider) *Repository {
	return &Repository{
		db: p,
		sb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
