package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
	"github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) AcceptFriendRequest(ctx context.Context, request *social.AcceptFriendRequestRequest) (*social.AcceptFriendRequestResponse, error) {

	data := &dto.AcceptFriendRequestIN{RequestID: request.RequestId}

	outUC, err := s.SocialUseCase.AcceptFriendRequest(ctx, data)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	friendRequest := &social.FriendRequest{
		RequestId: outUC.RequestID,
		Status:    social.Status(outUC.Status),
	}

	return &social.AcceptFriendRequestResponse{FriendRequest: friendRequest}, nil
}
