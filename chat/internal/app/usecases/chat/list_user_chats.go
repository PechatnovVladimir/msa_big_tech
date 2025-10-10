package chat

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
)

func (s *Service) ListUserChats(ctx context.Context, in *dto.ListUserChatsIN) (*dto.ListUserChatsOUT, error) {

	data := fromListUserChatsIN(in)

	chats, err := s.ChatRepo.ListUserChats(ctx, data)
	if err != nil {
		return nil, dto.ErrChatNotFound
	}

	out := toListUserChatsOUT(chats)

	return out, nil
}
