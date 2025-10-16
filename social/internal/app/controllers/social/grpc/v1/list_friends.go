package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
)

func (s *Service) ListFriends(ctx context.Context, request *social.ListFriendsRequest) (*social.ListFriendsResponse, error) {

	data := fromListFriendsRequestToDto(request)

	listFriendsWithCursor, err := s.SocialUseCase.ListFriends(ctx, data)
	if err != nil {
		return nil, err
	}

	out := fromDtoToListFriendsResponse(listFriendsWithCursor)

	return out, nil
}
