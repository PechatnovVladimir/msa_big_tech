package outbox

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social"
)

var _ social.OutboxRepository = (*Processor)(nil)

// Repository - репозиторий Outbox
type Repository interface {
	SaveEvent(ctx context.Context, e *Event) error
}

type TransactionManager interface {
	RunRepeatableRead(ctx context.Context, f func(ctx context.Context) error) error
}

type Deps struct {
	Repository Repository
}

type Processor struct {
	Deps
}

func NewProcessor(deps Deps) *Processor {
	return &Processor{
		Deps: deps,
	}
}
