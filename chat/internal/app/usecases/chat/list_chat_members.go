package chat

import (
	"context"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"slices"
)

func (s *Service) ListChatMembers(ctx context.Context, in *dto.ListChatMembersIN) (*dto.ListChatMembersOUT, error) {
	const api = "ChatService.ListChatMembers"

	currentUser, err := s.UserProvider.GetUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s. %s: %w", "Пользователь не аутентифицирован", api, chat.ErrUnauthenticated)
	}

	data := fromListChatMembersIN(in)

	chatInfo, err := s.ChatRepo.GetChat(ctx, data)

	if err != nil {
		return &dto.ListChatMembersOUT{}, fmt.Errorf("%s: %w", api, err)
	}

	userIDs := chatInfo.Members

	//если текущий пользователь не является участником чата, то доступ к чату запрещен
	if !slices.Contains(userIDs, currentUser.UserID) {
		return nil, fmt.Errorf("%s: %w", api, chat.ErrPermissionDenied)
	}

	out := toListChatMembersOUT(userIDs)

	return out, nil
}
