package chat

import (
	"context"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
)

func (s *Service) ListUserChats(ctx context.Context, in *dto.ListUserChatsIN) (*dto.ListUserChatsOUT, error) {
	const api = "ChatService.ListUserChats"

	currentUser, err := s.UserProvider.GetUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s. %s: %w", "Пользователь не аутентифицирован", api, chat.ErrUnauthenticated)
	}

	if currentUser.UserID != in.UserID {
		return nil, fmt.Errorf("%s: %w", api, chat.ErrPermissionDenied)
	}

	data := fromListUserChatsIN(in)

	chats, err := s.ChatRepo.ListUserChats(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	out := toListUserChatsOUT(chats)

	return out, nil
}
