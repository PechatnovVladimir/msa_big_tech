package outbox

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
	"github.com/google/uuid"
	"time"
)

func (p *Processor) SaveFriendRequest(ctx context.Context, friendRequest *social.FriendRequest) error {
	event := Event{
		ID:            uuid.New().String(),
		AggregateType: AggregateTypeFriendRequest,
		AggregateID:   friendRequest.RequestID,
		EventType:     EventTypeFriendRequest,
		Payload:       nil,
		CreatedAt:     time.Now().UTC(),
	}
	return p.Repository.SaveEvent(ctx, &event)
}
