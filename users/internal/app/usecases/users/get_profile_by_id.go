package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
)

func (s *Service) GetProfileByID(ctx context.Context, profileID string) (*users.UserProfile, error) {
	userProfile, err := s.repository.GetProfileByID(ctx, profileID)
	if err != nil {
		return nil, err
	}
	return userProfile, nil
}
