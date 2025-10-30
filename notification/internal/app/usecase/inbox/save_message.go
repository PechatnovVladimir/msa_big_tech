package inbox

import (
	"context"
	"github.com/IBM/sarama"
)

func (s *Service) SaveMessage(ctx context.Context, msg *sarama.ConsumerMessage) error {

	data := fromConsumerMessage(msg)

	err := s.InboxRepo.SaveMessage(ctx, data)
	if err != nil {
		return err
	}

	return nil
}
