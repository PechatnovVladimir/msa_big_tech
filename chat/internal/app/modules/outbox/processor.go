package outbox

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat"
)

var _ chat.OutboxRepository = (*Processor)(nil)

// Repository - репозиторий Outbox
type Repository interface {
	SaveEvent(ctx context.Context, e *Event) error
	SearchEvents(ctx context.Context, opts ...SearchEventsOption) []*Event
	UpdateEvents(ctx context.Context, opts ...UpdateEventsOption) error
}

type TransactionManager interface {
	RunRepeatableRead(ctx context.Context, f func(ctx context.Context) error) error
	//RunReadCommitted(ctx context.Context, f func(ctx context.Context) error) error
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
