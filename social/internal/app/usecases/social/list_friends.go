package social

import (
	"context"
	"fmt"
	models "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
)

func (s *Service) ListFriends(ctx context.Context, in *dto.ListFriendsIN) (*dto.ListFriendsOUT, error) {
	const api = "SocialService.ListFriends"

	_, err := s.UserProvider.GetUserFromContext(ctx)
	if err != nil {
		return &dto.ListFriendsOUT{}, fmt.Errorf("%s: %w", api, models.ErrSocialUnauthenticated)
	}

	userID, paginationOpts := fromListFriendsIN(in)

	userIDs, err := s.SocialRepo.ListFriends(ctx, userID, paginationOpts)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	out := toListFriendsOUT(userIDs)

	return out, nil

}
