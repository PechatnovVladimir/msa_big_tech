package v1

import (
	"buf.build/go/protovalidate"
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) SendMessage(ctx context.Context, request *chat.SendMessageRequest) (*chat.SendMessageResponse, error) {

	//валидация по proto описанию
	err := protovalidate.Validate(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	data := fromSendMessageRequestToDto(request)

	message, err := s.ChatUseCase.SendMessage(ctx, data)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	out := fromDtoToSendMessageResponse(message)

	return out, nil
}
