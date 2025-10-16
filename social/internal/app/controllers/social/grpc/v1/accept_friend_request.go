package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
	"github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
)

func (s *Service) AcceptFriendRequest(ctx context.Context, request *social.AcceptFriendRequestRequest) (*social.AcceptFriendRequestResponse, error) {

	data := &dto.AcceptFriendRequestIN{RequestID: request.RequestId}

	outUC, err := s.SocialUseCase.AcceptFriendRequest(ctx, data)
	if err != nil {
		return nil, err
	}

	friendRequest := &social.FriendRequest{
		RequestId: outUC.RequestID,
		Status:    social.Status(outUC.Status),
	}

	return &social.AcceptFriendRequestResponse{FriendRequest: friendRequest}, nil
}
