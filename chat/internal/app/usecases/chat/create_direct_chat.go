package chat

import (
	"context"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
)

// CreateDirectChat - создание личного чата
func (s *Service) CreateDirectChat(ctx context.Context, in *dto.CreateDirectChatIN) (*dto.CreateDirectChatOUT, error) {
	const api = "ChatService.CreateDirectChat"

	currentUser, err := s.UserProvider.GetUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s. %s: %w", "Пользователь не аутентифицирован", api, chat.ErrUnauthenticated)
	}

	if currentUser.UserID == in.ParticipantID {
		return nil, fmt.Errorf("%s. %s: %w", "Пользователь не может открыть чат сам с собой", api, chat.ErrInvalidArgument)
	}

	participantID := fromCreateDirectChatIN(in)

	//запись в репозиторий
	chatData, err := s.ChatRepo.CreateDirectChat(ctx, currentUser.UserID, participantID)
	if err != nil {
		return &dto.CreateDirectChatOUT{}, fmt.Errorf("%s. %s: %w", "Сбой записи в БД", api, err)
	}

	out := toCreateDirectChatOUT(chatData)
	//возвращаем ID созданного чата
	return out, nil
}
