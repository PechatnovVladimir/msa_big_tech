package chat

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"log"
)

func (s *Service) GetChat(ctx context.Context, in dto.GetChatIN) (dto.GetChatOut, error) {
	log.Println("chat usecase called GetChat")
	return dto.GetChatOut{}, nil
}
