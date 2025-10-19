package inbox

import (
	"context"
	"github.com/IBM/sarama"
)

func (r *Repository) SaveMessage(ctx context.Context, msg *sarama.ConsumerMessage) error {

	msg.Value
	return nil
}
