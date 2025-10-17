package social

import (
	"context"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
	models "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
)

func (s *Service) SendFriendRequest(ctx context.Context, in *dto.SendFriendRequestIN) (*dto.SendFriendRequestOUT, error) {
	const api = "SocialService.SendFriendRequest"

	currentUser, err := s.UserProvider.GetUserFromContext(ctx)
	if err != nil {
		return &dto.SendFriendRequestOUT{}, fmt.Errorf("%s: %w", api, models.ErrSocialUnauthenticated)
	}

	friendRequest := social.NewFriendRequest(currentUser.UserID, in.UserId)

	var outFriendRequest *social.FriendRequest

	//начало транзакции, в транзакции две операции - запись в основной репо и в репо outbox
	err = s.TransactionManager.RunReadCommitted(ctx, func(ctx context.Context) error {
		outFriendRequest, err = s.SocialRepo.SendFriendRequest(ctx, friendRequest)
		if err != nil {
			return err
		}
		err = s.OutboxRepo.SaveFriendRequest(ctx, outFriendRequest)
		if err != nil {
			return err
		}
		return nil
	})
	//конец транзакции

	if err != nil {
		return &dto.SendFriendRequestOUT{}, fmt.Errorf("%s: %w", api, err)
	}

	return &dto.SendFriendRequestOUT{
		RequestID: outFriendRequest.RequestID,
		Status:    dto.StatusRequest(outFriendRequest.Status),
	}, nil
}
