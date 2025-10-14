package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
)

func (s *Service) ListChatMembers(ctx context.Context, request *chat.ListChatMembersRequest) (*chat.ListChatMembersResponse, error) {
	data := fromListChatMembersRequestToDto(request)

	userIDs, err := s.ChatUseCase.ListChatMembers(ctx, data)

	if err != nil {
		return nil, err
	}

	out := fromDtoToListChatMembersResponse(userIDs)

	return out, nil
}
