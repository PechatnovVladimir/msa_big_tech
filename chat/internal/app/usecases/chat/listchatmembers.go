package chat

import (
	"context"
	dtoRepo "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/repositories/chat/dto"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"log"
)

func (s *Service) ListChatMembers(ctx context.Context, in dto.ListChatMembersIN) (dto.ListChatMembersOUT, error) {
	log.Println("chat usecase called  ListChatMembers")

	outRepo, err := s.ChatRepo.ListChatMembers(ctx, dtoRepo.ListChatMembersIN{})
	if err != nil {
		return dto.ListChatMembersOUT{}, err
	}

	_ = outRepo

	return dto.ListChatMembersOUT{}, nil
}
