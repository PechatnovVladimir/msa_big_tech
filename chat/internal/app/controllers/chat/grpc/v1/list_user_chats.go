package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) ListUserChats(ctx context.Context, request *chat.ListUserChatsRequest) (*chat.ListUserChatsResponse, error) {
	data := fromListUserChatsRequestToDto(request)

	chats, err := s.ChatUseCase.ListUserChats(ctx, data)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	out := fromDtoToListUserChatsResponse(chats)

	return out, nil
}
