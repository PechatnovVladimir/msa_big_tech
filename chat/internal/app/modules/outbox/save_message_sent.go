package outbox

import (
	"context"
	"encoding/json"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"github.com/google/uuid"
	"time"
)

func (p *Processor) SaveMessageSent(ctx context.Context, message *chat.Message) error {
	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	event := Event{
		ID:            uuid.New().String(),
		AggregateType: AggregateTypeMessageSent,
		AggregateID:   message.MessageID,
		EventType:     EventTypeMessageSent,
		Payload:       data,
		CreatedAt:     time.Now().UTC(),
	}
	return p.Repository.SaveEvent(ctx, &event)
}
