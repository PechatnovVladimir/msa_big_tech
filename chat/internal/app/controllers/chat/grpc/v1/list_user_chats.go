package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
)

func (s *Service) ListUserChats(ctx context.Context, request *chat.ListUserChatsRequest) (*chat.ListUserChatsResponse, error) {
	data := fromListUserChatsRequestToDto(request)

	chats, err := s.ChatUseCase.ListUserChats(ctx, data)

	if err != nil {
		return nil, err
	}

	out := fromDtoToListUserChatsResponse(chats)

	return out, nil
}
