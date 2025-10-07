package chat

import (
	"context"
	dtoRepo "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/repositories/chat/dto"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"log"
)

func (s *Service) ListMessages(ctx context.Context, in dto.ListMessagesIN) (dto.ListMessagesOUT, error) {
	log.Println("chat usecase called ListMessages")

	outRepo, err := s.ChatRepo.ListMessages(ctx, dtoRepo.ListMessagesIN{})
	if err != nil {
		return dto.ListMessagesOUT{}, err
	}
	_ = outRepo

	return dto.ListMessagesOUT{}, nil
}
