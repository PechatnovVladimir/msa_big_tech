package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
)

func (s *Service) GetChat(ctx context.Context, request *chat.GetChatRequest) (*chat.GetChatResponse, error) {
	data := fromGetChatRequestToDto(request)

	chatInfo, err := s.ChatUseCase.GetChat(ctx, data)

	if err != nil {
		return &chat.GetChatResponse{}, err
	}

	out := fromDtoToGetChatResponse(chatInfo)

	return out, nil
}
