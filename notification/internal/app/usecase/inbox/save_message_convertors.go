package inbox

import (
	"github.com/IBM/sarama"
	"github.com/PechatnovVladimir/msa_big_tech/notification/internal/app/model"
	"time"
)

func fromConsumerMessage(msg *sarama.ConsumerMessage) *model.Inbox {
	return &model.Inbox{
		ID:          string(msg.Key),
		Topic:       msg.Topic,
		Partition:   msg.Partition,
		Offset:      msg.Offset,
		Payload:     msg.Value,
		Status:      "received",
		Attempts:    0,
		LastError:   "",
		ReceivedAt:  time.Now(),
		ProcessedAt: nil,
	}
}
