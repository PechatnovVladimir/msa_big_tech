package chat

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"log"
)

func (s *Service) ListChatMembers(ctx context.Context, in dto.ListChatMembersIN) (dto.ListChatMembersOUT, error) {
	log.Println("chat usecase called  ListChatMembers")
	return dto.ListChatMembersOUT{}, nil
}
