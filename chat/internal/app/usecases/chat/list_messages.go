package chat

import (
	"context"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"slices"
)

func (s *Service) ListMessages(ctx context.Context, in *dto.ListMessagesIN) (*dto.ListMessagesOUT, error) {
	const api = "ChatService.ListMessages"

	currentUser, err := s.UserProvider.GetUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s. %s: %w", "Пользователь не аутентифицирован", api, chat.ErrUnauthenticated)
	}

	chatData, err := s.ChatRepo.GetChat(ctx, in.ChatID)

	if err != nil {
		return &dto.ListMessagesOUT{}, fmt.Errorf("%s: %w", api, err)
	}

	//если текущий пользователь не является участником чата, то доступ к чату запрещен
	if !slices.Contains(chatData.Members, currentUser.UserID) {
		return nil, fmt.Errorf("%s: %w", api, chat.ErrPermissionDenied)
	}

	chatID, paginationOpts := fromListMessagesIN(in)

	messages, err := s.ChatRepo.ListMessages(ctx, chatID, paginationOpts)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	out := toListMessageOUT(messages)

	return out, nil
}
