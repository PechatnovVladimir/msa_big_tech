package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
)

func (s *Service) UpdateProfile(ctx context.Context, d dto.UpdateProfileDTO) (*users.UserProfile, error) {
	_, err := s.repository.GetProfileByID(ctx, d.ID)
	if err != nil {
		return nil, users.ErrUserNotFound
	}

	userProfile := &users.UserProfileForUpdate{
		ID: d.ID,
	}

	if d.Email != nil {
		userProfile.Email = d.Email
	}

	if d.Nickname != nil {
		userProfile.Nickname = d.Nickname
	}

	if d.Bio != nil {
		userProfile.Bio = d.Bio
	}

	if d.Avatar != nil {
		userProfile.Avatar = d.Avatar
	}

	out, err := s.repository.UpdateProfile(ctx, userProfile)
	if err != nil {
		return nil, err
	}

	return out, nil
}
