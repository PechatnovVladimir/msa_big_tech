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
	_ = currentUser

	data := fromGetChatIN(in)

	//получаем информацию о чате
	chatData, err := s.ChatRepo.GetChat(ctx, data)

	if err != nil {
		return &dto.GetChatOUT{}, dto.ErrChatNotFound
	}

	chatMembers, _ := s.ListChatMembers(ctx, &dto.ListChatMembersIN{ChatID: chatData.ChatID})

	if !slices.Contains(convertToStringSlice(chatMembers.UserIDs), currentUser.UserID) {
		return nil, fmt.Errorf("%s: %w", api, chat.ErrPermissionDenied)
	}

	return toGetChatOUT(chatData), nil
}

func convertToStringSlice(s []*string) []string {
	s1 := make([]string, 0, len(s))
	for _, ptr := range s {
		if ptr != nil {
			s1 = append(s1, *ptr)
		} else {
			s1 = append(s1, "")
		}
	}
	return s1
}
