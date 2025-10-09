package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
)

func (s *Service) GetProfileByNickname(ctx context.Context, in dto.GetProfileByNickname) (*users.UserProfile, error) {

	userProfile, err := s.repository.GetProfileByNickname(ctx, in.Nickname)
	if err != nil {
		return nil, users.ErrUserNotFound
	}

	return userProfile, nil
}
