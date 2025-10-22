package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
)

func (s *Service) SendMessage(ctx context.Context, request *chat.SendMessageRequest) (*chat.SendMessageResponse, error) {
	data := fromSendMessageRequestToDto(request)

	message, err := s.ChatUseCase.SendMessage(ctx, data)

	if err != nil {
		return nil, err
	}

	out := fromDtoToSendMessageResponse(message)

	return out, nil
}
