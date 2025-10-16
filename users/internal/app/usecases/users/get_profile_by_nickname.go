package users

import (
	"context"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
)

func (s *Service) GetProfileByNickname(ctx context.Context, in dto.GetProfileByNickname) (*users.UserProfile, error) {
	const api = "UserService.GetProfileByNickname"

	userProfile, err := s.UserRepo.GetProfileByNickname(ctx, in.Nickname)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, users.ErrUserNotFound)
	}

	return userProfile, nil
}
