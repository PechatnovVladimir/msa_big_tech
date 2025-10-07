package users

import (
	"github.com/Masterminds/squirrel"
	"github.com/PechatnovVladimir/msa_big_tech/users/pkg/postgres/users"
)

type Repository struct {
	db users.TransactionManagerAPI
	sb squirrel.StatementBuilderType
}

func NewRepository(txManager users.TransactionManagerAPI) *Repository {
	return &Repository{
		db: txManager,
		sb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
