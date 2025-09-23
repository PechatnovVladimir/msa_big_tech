package chat

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"log"
)

func (s *Service) StreamMessages(ctx context.Context, in dto.StreamMessagesIN) (dto.StreamMessagesOUT, error) {
	log.Println("chat usecase called  StreamMessages")
	return dto.StreamMessagesOUT{}, nil
}
