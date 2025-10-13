package social

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
)

func (s *Service) SendFriendRequest(ctx context.Context, in *dto.SendFriendRequestIN) (*dto.SendFriendRequestOUT, error) {
	currentUser, err := s.UserProvider.GetUserFromContext(ctx)
	if err != nil {
		return &dto.SendFriendRequestOUT{}, err
	}

	friendRequest := social.NewFriendRequest(currentUser.UserID, in.UserId)

	outFriendRequest, err := s.SocialRepo.SendFriendRequest(ctx, friendRequest)

	if err != nil {
		return &dto.SendFriendRequestOUT{}, err
	}

	return &dto.SendFriendRequestOUT{
		RequestID: outFriendRequest.RequestID,
		Status:    dto.StatusRequest(outFriendRequest.Status),
	}, nil
}
