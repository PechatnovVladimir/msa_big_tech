package v1

import (
	"buf.build/go/protovalidate"
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetChat(ctx context.Context, request *chat.GetChatRequest) (*chat.GetChatResponse, error) {
	//валидация по proto описанию
	err := protovalidate.Validate(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	data := fromGetChatRequestToDto(request)

	chatInfo, err := s.ChatUseCase.GetChat(ctx, data)

	if err != nil {
		return &chat.GetChatResponse{}, dto.ErrChatNotFound
	}

	out := fromDtoToGetChatResponse(chatInfo)

	return out, nil
}
