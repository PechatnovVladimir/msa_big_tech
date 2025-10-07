package chat

import (
	"context"
	dtoRepo "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/repositories/chat/dto"
)

func (r *Repository) GetChatByUserAndParticipant(ctx context.Context, in dtoRepo.GetChatByUserAndParticipantIN) (out dtoRepo.GetChatByUserAndParticipantOUT, ok bool) {
	return dtoRepo.GetChatByUserAndParticipantOUT{}, false
}
