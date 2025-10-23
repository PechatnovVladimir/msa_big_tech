package users

import (
	"github.com/Masterminds/squirrel"
	"github.com/PechatnovVladimir/msa_big_tech/lib/postgres"
)

type Repository struct {
	db postgres.QueryEngineProvider
	sb squirrel.StatementBuilderType
}

func NewRepository(txManager postgres.QueryEngineProvider) *Repository {
	return &Repository{
		db: txManager,
		sb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
