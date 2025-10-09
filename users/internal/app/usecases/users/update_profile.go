package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
)

func (s *Service) UpdateProfile(ctx context.Context, in dto.UpdateProfile) (*users.UserProfile, error) {
	_, err := s.repository.GetProfileByID(ctx, in.ID)
	if err != nil {
		return nil, users.ErrUserNotFound
	}

	data := modelUserProfileForUpdateFromUpdateProfileDto(in)

	updatedProfile, err := s.repository.UpdateProfile(ctx, data)
	if err != nil {
		return nil, err
	}

	return updatedProfile, nil
}
