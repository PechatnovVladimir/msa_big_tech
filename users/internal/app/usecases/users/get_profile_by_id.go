package users

import (
	"context"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
)

func (s *Service) GetProfileByID(ctx context.Context, in dto.GetProfileById) (*users.UserProfile, error) {
	const api = "UserService.GetProfileByID"

	userProfile, err := s.UserRepo.GetProfileByID(ctx, in.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, users.ErrUserNotFound)
	}
	return userProfile, nil
}
