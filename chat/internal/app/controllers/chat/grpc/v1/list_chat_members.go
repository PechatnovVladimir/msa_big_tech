package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) ListChatMembers(ctx context.Context, request *chat.ListChatMembersRequest) (*chat.ListChatMembersResponse, error) {
	data := fromListChatMembersRequestToDto(request)

	userIDs, err := s.ChatUseCase.ListChatMembers(ctx, data)

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	out := fromDtoToListChatMembersResponse(userIDs)

	return out, nil
}
