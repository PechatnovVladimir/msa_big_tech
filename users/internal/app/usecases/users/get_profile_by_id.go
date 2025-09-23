package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
)

func (s *UserService) GetProfileByID(ctx context.Context, profileID string) (*users.UserProfile, error) {
	userProfile, err := s.userRepo.GetProfileByID(ctx, profileID)
	if err != nil {
		return nil, err
	}
	return userProfile, nil
}
