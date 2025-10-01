package chat

import (
	"context"
	dtoRepo "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/repositories/chat/dto"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"log"
)

func (s *Service) ListUserChats(ctx context.Context, in dto.ListUserChatsIN) (dto.ListUserChatsOUT, error) {
	log.Println("chat usecase called  ListUserChats called")
	outRepo, err := s.ChatRepo.ListUserChats(ctx, dtoRepo.ListUserChatsIN{})
	if err != nil {
		return dto.ListUserChatsOUT{}, err
	}

	_ = outRepo

	return dto.ListUserChatsOUT{}, nil
}
