package outbox

import (
	"context"
	"github.com/google/uuid"
	"time"
)

func (p *Processor) SaveFriendRequestID(ctx context.Context, friendRequestID string) error {
	event := Event{
		ID:            uuid.New(),
		AggregateType: AggregateTypeFriendRequest,
		AggregateID:   friendRequestID,
		EventType:     EventTypeFriendRequest,
		Payload:       nil,
		CreatedAt:     time.Now().UTC(),
	}
	return p.Repository.SaveEvent(ctx, &event)
}
