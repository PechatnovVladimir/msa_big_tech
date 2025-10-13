package social

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/converter"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
)

func (s *Service) SendFriendRequest(ctx context.Context, in dto.SendFriendRequestIN) (dto.SendFriendRequestOUT, error) {

	authUser, err := s.AuthProvider.GetAuthUser()
	if err != nil {
		return dto.SendFriendRequestOUT{}, err
	}

	friendRequest := models.NewFriendRequest(authUser, in.UserId)

	inRepo, err := converter.SendFriendRequestDtoToRepo(ctx, friendRequest)
	if err != nil {
		return dto.SendFriendRequestOUT{}, err
	}

	err = s.SocialRepo.SendFriendRequest(ctx, inRepo)
	if err != nil {
		return dto.SendFriendRequestOUT{}, err
	}

	return dto.SendFriendRequestOUT{
		RequestID: friendRequest.RequestID,
		Status:    dto.StatusRequest(friendRequest.Status),
	}, nil
}
