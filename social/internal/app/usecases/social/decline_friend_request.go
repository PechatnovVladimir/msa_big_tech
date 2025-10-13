package social

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
)

func (s *Service) DeclineFriendRequest(ctx context.Context, in *dto.DeclineFriendRequestIN) (*dto.DeclineFriendRequestOUT, error) {

	friendRequest, err := s.SocialRepo.ChangeStatusFriendRequest(ctx, in.RequestID, int(dto.DECLINED))

	if err != nil {
		return &dto.DeclineFriendRequestOUT{}, err
	}

	return &dto.DeclineFriendRequestOUT{
		RequestID: friendRequest.RequestID,
		Status:    dto.StatusRequest(friendRequest.Status),
	}, nil
}
