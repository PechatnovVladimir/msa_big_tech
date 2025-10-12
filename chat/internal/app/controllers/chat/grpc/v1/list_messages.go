package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) ListMessages(ctx context.Context, request *chat.ListMessagesRequest) (*chat.ListMessagesResponse, error) {
	data := fromListMessagesRequest(request)

	messagesWithCursor, err := s.ChatUseCase.ListMessages(ctx, data)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	out := fromDtoToListMessagesResponse(messagesWithCursor)

	return out, nil
}
