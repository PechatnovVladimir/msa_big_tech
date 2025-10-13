package social

import (
	"github.com/Masterminds/squirrel"
	"github.com/PechatnovVladimir/msa_big_tech/pkg/postgres"
)

type Repository struct {
	db postgres.TransactionManagerAPI
	sb squirrel.StatementBuilderType
}

func New(txManager postgres.TransactionManagerAPI) *Repository {
	return &Repository{
		db: txManager,
		sb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
