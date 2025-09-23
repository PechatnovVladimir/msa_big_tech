package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
)

func (s *Service) UpdateProfile(ctx context.Context, d dto.UpdateProfileDTO) (*users.UserProfile, error) {
	userProfile, err := s.repository.GetProfileByID(ctx, d.ID)
	if err != nil {
		return nil, users.ErrUserNotFound
	}

	if d.Bio != "" {
		userProfile.Bio = d.Bio
	}

	if d.Avatar != "" {
		userProfile.Avatar = d.Avatar
	}

	err = s.repository.UpdateProfile(ctx, userProfile)
	if err != nil {
		return nil, err
	}

	return userProfile, nil
}
