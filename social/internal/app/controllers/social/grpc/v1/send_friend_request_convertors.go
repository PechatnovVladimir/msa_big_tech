package v1

import (
	models "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
	"github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
)

func fromSendFriendRequestRequestToDto(in *social.SendFriendRequestRequest) (*dto.SendFriendRequestIN, error) {
	if in == nil {
		return &dto.SendFriendRequestIN{}, models.ErrSocialInvalidArgument
	}

	out := &dto.SendFriendRequestIN{
		UserId: in.UserId,
	}
	return out, nil
}

func fromDtoToSendFriendRequestResponse(in *dto.SendFriendRequestOUT) (*social.SendFriendRequestResponse, error) {
	if in.RequestID == "" {
		return nil, models.ErrSocialInvalidArgument
	}

	out := &social.SendFriendRequestResponse{
		FriendRequest: &social.FriendRequest{
			RequestId: in.RequestID,
			Status:    social.Status(in.Status),
		},
	}

	return out, nil
}
