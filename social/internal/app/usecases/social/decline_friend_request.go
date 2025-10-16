package social

import (
	"context"
	"fmt"
	models "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
)

func (s *Service) DeclineFriendRequest(ctx context.Context, in *dto.DeclineFriendRequestIN) (*dto.DeclineFriendRequestOUT, error) {
	const api = "SocialService.AcceptFriendRequest"

	currentUser, err := s.UserProvider.GetUserFromContext(ctx)
	if err != nil {
		return &dto.DeclineFriendRequestOUT{}, fmt.Errorf("%s: %w", api, models.ErrSocialUnauthenticated)
	}

	request, err := s.SocialRepo.GetFriendRequest(ctx, in.RequestID)
	if err != nil {
		return &dto.DeclineFriendRequestOUT{}, fmt.Errorf("%s: %w", api, err)
	}

	if request == nil {
		return &dto.DeclineFriendRequestOUT{}, fmt.Errorf("%s: %w", api, models.ErrSocialNotFound)
	}

	//если заявка не текущему пользователю, то нельзя...
	if request.ToUserID != currentUser.UserID {
		return &dto.DeclineFriendRequestOUT{}, fmt.Errorf("%s: %w", api, models.ErrSocialPermissionDenied)
	}

	friendRequest, err := s.SocialRepo.ChangeStatusFriendRequest(ctx, in.RequestID, int(models.DECLINED))
	if err != nil {
		return &dto.DeclineFriendRequestOUT{}, err
	}

	return &dto.DeclineFriendRequestOUT{
		RequestID: friendRequest.RequestID,
		Status:    dto.StatusRequest(friendRequest.Status),
	}, nil
}
