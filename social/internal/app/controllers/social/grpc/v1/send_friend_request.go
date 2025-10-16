package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
)

func (s *Service) SendFriendRequest(ctx context.Context, request *social.SendFriendRequestRequest) (*social.SendFriendRequestResponse, error) {

	data, err := fromSendFriendRequestRequestToDto(request)
	if err != nil {
		return nil, err
	}

	friendRequest, err := s.SocialUseCase.SendFriendRequest(ctx, data)
	if err != nil {
		return nil, err
	}

	out, err := fromDtoToSendFriendRequestResponse(friendRequest)
	if err != nil {
		return nil, err
	}

	return out, nil
}
