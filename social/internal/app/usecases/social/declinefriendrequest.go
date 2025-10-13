package social

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/converter"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
	"log"
)

func (s *Service) DeclineFriendRequest(ctx context.Context, in dto.DeclineFriendRequestIN) (dto.DeclineFriendRequestOUT, error) {
	log.Println("DeclineFriendRequest")
	authUser, err := s.AuthProvider.GetAuthUser()
	if err != nil {
		return dto.DeclineFriendRequestOUT{}, err
	}

	friendRequestFromRepo, err := s.SocialRepo.GetFriendRequestByID(ctx, in.RequestID)
	if err != nil {
		return dto.DeclineFriendRequestOUT{}, err
	}

	friendRequest, err := converter.FriendRequestsFromRepoToModel(ctx, friendRequestFromRepo)
	if err != nil {
		return dto.DeclineFriendRequestOUT{}, err
	}

	if friendRequest.ToUserID != authUser {
		return dto.DeclineFriendRequestOUT{}, models.ErrSocialPermissionDenied
	}

	friendRequest.Decline()

	friendRequestToRepo, err := converter.FriendRequestFromModelToRepo(ctx, friendRequest)
	if err != nil {
		return dto.DeclineFriendRequestOUT{}, err
	}

	err = s.SocialRepo.UpdateFriendRequest(ctx, friendRequestToRepo)
	if err != nil {
		return dto.DeclineFriendRequestOUT{}, err
	}

	return dto.DeclineFriendRequestOUT{
		RequestID: friendRequest.RequestID,
		Status:    dto.StatusRequest(friendRequest.Status),
	}, nil
}
