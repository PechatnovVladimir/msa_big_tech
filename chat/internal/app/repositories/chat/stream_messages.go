package chat

import (
	"context"
	"errors"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
)

func (r *Repository) StreamMessage(ctx context.Context, chatID string) (<-chan *chat.Message, error) {
	return nil, errors.New("not implemented")
}
