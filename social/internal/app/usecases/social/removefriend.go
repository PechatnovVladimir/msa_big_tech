package social

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models"
	dtoRepo "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/repositories/social/dto"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
	"log"
)

func (s *Service) RemoveFriend(ctx context.Context, in dto.RemoveFriendIN) error {
	log.Println("RemoveFriend")

	authUser, err := s.AuthProvider.GetAuthUser()
	if err != nil {
		return err
	}

	err = s.SocialRepo.DeleteFriendRequest(ctx, dtoRepo.DeleteFriendRequestIN{
		FromUserID: authUser,
		ToUserID:   in.UserID,
	})

	if err != nil {
		return models.ErrSocialNotFound
	}

	return nil
}
