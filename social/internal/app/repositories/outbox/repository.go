package outbox

import (
	"github.com/Masterminds/squirrel"
	"github.com/PechatnovVladimir/msa_big_tech/pkg/postgres"
)

type Repository struct {
	db postgres.TransactionManagerAPI
	qb squirrel.StatementBuilderType
}

// NewRepository конструктор Repository
func NewRepository(p postgres.TransactionManagerAPI) *Repository {
	return &Repository{
		db: p,
		qb: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
