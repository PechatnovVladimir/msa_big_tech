package inbox

import (
	"context"
	"encoding/json"
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

		var data model.Message

		if message.Topic == "chat.message.sent" {
			err := json.Unmarshal([]byte(message.Payload), &data)
			if err != nil {
				log.Println("[error unmarshal message]", err)
			}
		}

		if i%5 == 0 {
			failed = append(failed, message.ID)
			log.Println("failed message", message.Topic, message.ID, data.Text)
		} else {
			succeeded = append(succeeded, message.ID)
			log.Println("succeeded message", message.Topic, message.ID, data.Text)
		}
	}
	return succeeded, failed, nil
}

func NewNotificator() *Notificator {
	return &Notificator{}
}
