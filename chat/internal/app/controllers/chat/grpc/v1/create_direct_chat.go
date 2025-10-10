package v1

import (
	"buf.build/go/protovalidate"
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateDirectChat - delivery создать личный чат
func (s *Service) CreateDirectChat(ctx context.Context, request *chat.CreateDirectChatRequest) (*chat.CreateDirectChatResponse, error) {
	//валидация по proto описанию
	err := protovalidate.Validate(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	data := fromCreateDirectChatRequestToDTO(request)

	chatInfo, err := s.ChatUseCase.CreateDirectChat(ctx, data)
	if err != nil {
		return nil, err
	}

	out := fromDtoToCreateDirectChatResponse(chatInfo)

	return out, nil
}
