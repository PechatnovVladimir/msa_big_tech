package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
)

func (s *Service) GetProfileByID(ctx context.Context, in dto.GetProfileById) (*users.UserProfile, error) {

	userProfile, err := s.repository.GetProfileByID(ctx, in.ID)
	if err != nil {
		return nil, users.ErrUserNotFound
	}
	return userProfile, nil
}
