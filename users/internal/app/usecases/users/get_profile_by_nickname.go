package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
)

func (s *Service) GetProfileByNickname(ctx context.Context, nickname string) (*users.UserProfile, error) {
	userProfile, err := s.repository.GetProfileByNickname(ctx, nickname)
	if err != nil {
		return nil, err
	}

	return userProfile, nil
}
