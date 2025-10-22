package social

import (
	"context"
	"fmt"
	models "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
)

func (s *Service) RemoveFriend(ctx context.Context, in dto.RemoveFriendIN) error {
	const api = "SocialService.RemoveFriend"

	currentUser, err := s.UserProvider.GetUserFromContext(ctx)
	if err != nil {
		return fmt.Errorf("%s: %w", api, models.ErrSocialUnauthenticated)
	}

	err = s.SocialRepo.RemoveFriend(ctx, currentUser.UserID, in.UserID)
	if err != nil {
		return fmt.Errorf("%s: %w", api, err)
	}

	return nil
}
