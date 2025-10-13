package social

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/converter"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
	"log"
)

func (s *Service) AcceptFriendRequest(ctx context.Context, in dto.AcceptFriendRequestIN) (dto.AcceptFriendRequestOUT, error) {
	log.Println("AcceptFriendRequest")

	authUser, err := s.AuthProvider.GetAuthUser()
	if err != nil {
		return dto.AcceptFriendRequestOUT{}, err
	}

	friendRequestFromRepo, err := s.SocialRepo.GetFriendRequestByID(ctx, in.RequestID)
	if err != nil {
		return dto.AcceptFriendRequestOUT{}, err
	}

	friendRequest, err := converter.FriendRequestsFromRepoToModel(ctx, friendRequestFromRepo)
	if err != nil {
		return dto.AcceptFriendRequestOUT{}, err
	}

	if friendRequest.ToUserID != authUser {
		return dto.AcceptFriendRequestOUT{}, social.ErrSocialPermissionDenied
	}

	friendRequest.Accept()

	friendRequestToRepo, err := converter.FriendRequestFromModelToRepo(ctx, friendRequest)
	if err != nil {
		return dto.AcceptFriendRequestOUT{}, err
	}

	err = s.SocialRepo.UpdateFriendRequest(ctx, friendRequestToRepo)
	if err != nil {
		return dto.AcceptFriendRequestOUT{}, err
	}

	return dto.AcceptFriendRequestOUT{
		RequestID: friendRequest.RequestID,
		Status:    dto.StatusRequest(friendRequest.Status),
	}, nil
}
