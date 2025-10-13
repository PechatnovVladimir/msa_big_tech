package social

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
)

func (s *Service) AcceptFriendRequest(ctx context.Context, in *dto.AcceptFriendRequestIN) (*dto.AcceptFriendRequestOUT, error) {

	friendRequest, err := s.SocialRepo.ChangeStatusFriendRequest(ctx, in.RequestID, int(dto.ACCEPTED))
	if err != nil {
		return &dto.AcceptFriendRequestOUT{}, err
	}

	err = s.SocialRepo.CreateFriend(ctx, friendRequest)
	if err != nil {
		return &dto.AcceptFriendRequestOUT{}, err
	}

	return &dto.AcceptFriendRequestOUT{
		RequestID: friendRequest.RequestID,
		Status:    dto.StatusRequest(friendRequest.Status),
	}, nil

}
