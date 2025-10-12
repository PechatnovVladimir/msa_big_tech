package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
)

// CreateDirectChat - delivery создать личный чат
func (s *Service) CreateDirectChat(ctx context.Context, request *chat.CreateDirectChatRequest) (*chat.CreateDirectChatResponse, error) {
	data := fromCreateDirectChatRequestToDTO(request)

	chatInfo, err := s.ChatUseCase.CreateDirectChat(ctx, data)
	if err != nil {
		return nil, err
	}

	out := fromDtoToCreateDirectChatResponse(chatInfo)

	return out, nil
}
