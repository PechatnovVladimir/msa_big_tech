package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
)

func (s *Service) ListMessages(ctx context.Context, request *chat.ListMessagesRequest) (*chat.ListMessagesResponse, error) {
	data := fromListMessagesRequest(request)

	messagesWithCursor, err := s.ChatUseCase.ListMessages(ctx, data)
	if err != nil {
		return nil, err
	}

	out := fromDtoToListMessagesResponse(messagesWithCursor)

	return out, nil
}
