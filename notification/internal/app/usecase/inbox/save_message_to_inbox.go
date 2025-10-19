package inbox

import (
	"context"
	"github.com/IBM/sarama"
)

func (s *Service) SaveMessage(ctx context.Context, msg *sarama.ConsumerMessage) error {

	err := s.InboxRepo.SaveMessage(ctx, msg)
	if err != nil {
		return err
	}

	return nil
}
