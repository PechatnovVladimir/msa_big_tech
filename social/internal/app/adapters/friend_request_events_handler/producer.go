package friend_request_events_handler

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/modules/outbox"
	"log"
	"time"
)

type KafkaFriendRequestBatchHandler struct {
	maxBatchSize int
	closeTimeout time.Duration
}

func NewKafkaFriendRequestBatchHandler() *KafkaFriendRequestBatchHandler {
	return &KafkaFriendRequestBatchHandler{
		maxBatchSize: 100,
		closeTimeout: 10 * time.Second,
	}
}

func (h *KafkaFriendRequestBatchHandler) HandleBatch(ctx context.Context, events []*outbox.Event) (succeeded []string, failed []string, err error) {
	succeeded = make([]string, 0, len(events))
	failed = make([]string, 0, len(events))

	for i, event := range events {
		log.Println("Событие: ", event.ID)
		if i%2 == 0 {
			succeeded = append(succeeded, event.ID)
		} else {
			failed = append(failed, event.ID)
		}
	}

	return succeeded, failed, nil
}
