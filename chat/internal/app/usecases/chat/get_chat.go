package chat

import (
	"context"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"slices"
)

func (s *Service) GetChat(ctx context.Context, in *dto.GetChatIN) (*dto.GetChatOUT, error) {
	const api = "ChatService.GetChat"

	currentUser, err := s.UserProvider.GetUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s. %s: %w", "Пользователь не аутентифицирован", api, chat.ErrUnauthenticated)
	}

	data := fromGetChatIN(in)

	//получаем информацию о чате
	chatData, err := s.ChatRepo.GetChat(ctx, data)

	if err != nil {
		return &dto.GetChatOUT{}, fmt.Errorf("%s: %w", api, err)
	}

	//если текущий пользователь не является участником чата, то доступ к чату запрещен
	if !slices.Contains(chatData.Members, currentUser.UserID) {
		return nil, fmt.Errorf("%s: %w", api, chat.ErrPermissionDenied)
	}

	return toGetChatOUT(chatData), nil
}
