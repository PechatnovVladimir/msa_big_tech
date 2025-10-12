package chat

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"github.com/google/uuid"
	"log"
	"time"
)

func (s *Service) SendMessage(ctx context.Context, in *dto.SendMessageIN) (*dto.SendMessageOUT, error) {
	log.Println("chat usecase called  SendMessage")

	data := fromSendMessageIN(in)

	data.MessageID = uuid.New().String()

	currenUser, err := s.UserProvider.GetUserFromContext(ctx)

	if err != nil {
		return &dto.SendMessageOUT{}, dto.ErrChatPermissionDenied
	}

	data.SenderID = currenUser.UserID
	data.CreatedAt = time.Now()

	message, err := s.ChatRepo.SendMessage(ctx, data)

	if err != nil {
		return &dto.SendMessageOUT{}, err
	}

	out := toSendMessageOUT(message)

	return out, nil
}
