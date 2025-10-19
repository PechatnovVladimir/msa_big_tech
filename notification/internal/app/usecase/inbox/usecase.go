package inbox

import (
	"context"
	"github.com/IBM/sarama"
)

type Repository interface {
	SaveMessage(context.Context, *sarama.ConsumerMessage) error
}

type Deps struct {
	InboxRepo Repository
}

type Service struct {
	Deps
}

type UseCase interface {
	SaveMessage(ctx context.Context, msg *sarama.ConsumerMessage) error
}

var _ UseCase = (*Service)(nil)

func New(d Deps) *Service {
	return &Service{
		Deps: d,
	}
}
