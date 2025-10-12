package chat

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
)

func (s *Service) ListMessages(ctx context.Context, in *dto.ListMessagesIN) (*dto.ListMessagesOUT, error) {
	chatID, paginationOpts := fromListMessagesIN(in)

	messages, err := s.ChatRepo.ListMessages(ctx, chatID, paginationOpts)

	if err != nil {
		return nil, dto.ErrChatNotFound
	}

	out := toListMessageOUT(messages)

	return out, nil
}
