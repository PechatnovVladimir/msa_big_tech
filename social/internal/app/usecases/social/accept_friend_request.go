package social

import (
	"context"
	"fmt"
	models "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
)

func (s *Service) AcceptFriendRequest(ctx context.Context, in *dto.AcceptFriendRequestIN) (*dto.AcceptFriendRequestOUT, error) {
	const api = "SocialService.AcceptFriendRequest"

	currentUser, err := s.UserProvider.GetUserFromContext(ctx)
	if err != nil {
		return &dto.AcceptFriendRequestOUT{}, fmt.Errorf("%s: %w", api, models.ErrSocialUnauthenticated)
	}

	request, err := s.SocialRepo.GetFriendRequest(ctx, in.RequestID)
	if err != nil {
		return &dto.AcceptFriendRequestOUT{}, fmt.Errorf("%s: %w", api, err)
	}

	if request == nil {
		return &dto.AcceptFriendRequestOUT{}, fmt.Errorf("%s: %w", api, models.ErrSocialNotFound)
	}

	//если заявка не текущему пользователю, то нельзя...
	if request.ToUserID != currentUser.UserID {
		return &dto.AcceptFriendRequestOUT{}, fmt.Errorf("%s: %w", api, models.ErrSocialPermissionDenied)
	}

	//старт транзакции
	var friendRequest *models.FriendRequest
	err = s.Tm.RunReadCommitted(ctx,
		func(txCtx context.Context) error {
			friendRequest, err = s.SocialRepo.ChangeStatusFriendRequest(txCtx, in.RequestID, int(models.ACCEPTED))
			if err != nil {
				return fmt.Errorf("%s: %w", api, err)
			}

			err = s.SocialRepo.CreateFriend(txCtx, friendRequest)
			if err != nil {
				return fmt.Errorf("%s: %w", api, err)
			}
			return nil
		},
	)
	if err != nil {
		return &dto.AcceptFriendRequestOUT{}, fmt.Errorf("%s: %w", api, err)
	}

	return &dto.AcceptFriendRequestOUT{
		RequestID: friendRequest.RequestID,
		Status:    dto.StatusRequest(friendRequest.Status),
	}, nil

}
