package inbox

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/notification/internal/app/model"
	"log"
	"time"
)

type Notificator struct {
}

func (n *Notificator) HandleBatch(ctx context.Context, messages []*model.Inbox) (succeeded []string, failed []string, err error) {
	time.Sleep(time.Second * 5)
	succeeded = make([]string, 0, len(messages))
	failed = make([]string, 0, len(messages))

	for i, message := range messages {
		if i%5 == 0 {
			failed = append(failed, message.ID)
			log.Println("failed message", message.ID)
		} else {
			succeeded = append(succeeded, message.ID)
			log.Println("succeeded message", message.ID)
		}
	}
	return succeeded, failed, nil
}

func NewNotificator() *Notificator {
	return &Notificator{}
}
