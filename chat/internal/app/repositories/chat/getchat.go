package chat

import (
	"context"
	dtoRepo "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/repositories/chat/dto"
	"github.com/google/uuid"
)

func (r *Repository) GetChat(ctx context.Context, in dtoRepo.GetChatIN) (out dtoRepo.GetChatOUT, err error) {
	//select в pg
	out = dtoRepo.GetChatOUT{
		ChatID:        in.ChatID,
		UserID:        uuid.New().String(),
		ParticipantID: uuid.New().String(),
	}
	return out, nil
}
