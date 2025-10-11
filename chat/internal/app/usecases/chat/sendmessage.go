package chat

import (
	"context"
	dtoRepo "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/repositories/chat/dto"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"log"
)

func (s *Service) SendMessage(ctx context.Context, in dto.SendMessageIN) (dto.SendMessageOut, error) {
	log.Println("chat usecase called  SendMessage")

	outRepo, err := s.ChatRepo.SendMessage(ctx, dtoRepo.SendMessageIN{})
	if err != nil {
		return dto.SendMessageOut{}, err
	}

	_ = outRepo

	return dto.SendMessageOut{}, nil
}
