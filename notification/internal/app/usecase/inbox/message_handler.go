package inbox

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/notification/internal/app/model"
	"log"
)

type Notificator struct {
}

func (n *Notificator) HandleBatch(ctx context.Context, messages []*model.Inbox) (succeeded []string, failed []string, err error) {
	succeeded = make([]string, 0, len(messages))
	failed = make([]string, 0, len(messages))

	for i, message := range messages {
		if i%2 == 0 {
			succeeded = append(succeeded, message.ID)
			log.Println("succeeded message", message.ID)
		} else {
			failed = append(failed, message.ID)
			log.Println("failed message", message.ID)
		}
	}
	return succeeded, failed, nil
}

func NewNotificator() *Notificator {
	return &Notificator{}
}
