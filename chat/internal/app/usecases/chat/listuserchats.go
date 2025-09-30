package chat

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"log"
)

func (s *Service) ListUserChats(ctx context.Context, in dto.ListUserChatsIN) (dto.ListUserChatsOUT, error) {
	log.Println("chat usecase called  ListUserChats called")
	return dto.ListUserChatsOUT{}, nil
}
