package chat

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"log"
)

func (s *Service) SendMessage(ctx context.Context, in dto.SendMessageIN) (dto.SendMessageOut, error) {
	log.Println("chat usecase called  SendMessage")
	return dto.SendMessageOut{}, nil
}
