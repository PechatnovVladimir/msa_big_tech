package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) SendFriendRequest(ctx context.Context, request *social.SendFriendRequestRequest) (*social.SendFriendRequestResponse, error) {

	data, err := fromSendFriendRequestRequestToDto(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	friendRequest, err := s.SocialUseCase.SendFriendRequest(ctx, data)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	out, err := fromDtoToSendFriendRequestResponse(friendRequest)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return out, nil
}
