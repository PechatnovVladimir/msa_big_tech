package outbox

import (
	"context"
	"encoding/json"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
	"github.com/google/uuid"
	"time"
)

func (p *Processor) SaveFriendUpdated(ctx context.Context, friendRequest *social.FriendRequest) error {
	data, err := json.Marshal(friendRequest)
	if err != nil {
		return err
	}

	event := Event{
		ID:            uuid.New().String(),
		AggregateType: AggregateTypeFriendUpdated,
		AggregateID:   friendRequest.RequestID,
		EventType:     EventTypeFriendUpdated,
		Payload:       data,
		CreatedAt:     time.Now().UTC(),
	}
	return p.Repository.SaveEvent(ctx, &event)
}
