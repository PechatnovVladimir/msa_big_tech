package chat

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"log"
)

func (s *Service) ListMessages(ctx context.Context, in dto.ListMessagesIN) (dto.ListMessagesOUT, error) {
	log.Println("chat usecase called ListMessages")
	return dto.ListMessagesOUT{}, nil
}
