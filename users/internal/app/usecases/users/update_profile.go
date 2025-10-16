package users

import (
	"context"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
)

func (s *Service) UpdateProfile(ctx context.Context, profile dto.UpdateProfile) (*users.UserProfile, error) {
	const api = "UserService.UpdateProfile"

	currentUser, err := s.UserProvider.GetUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, users.ErrUnauthenticated)
	}

	if currentUser != profile.ID {
		return nil, fmt.Errorf("%s: %w", api, users.ErrPermissionDenied)
	}

	_, err = s.UserRepo.GetProfileByID(ctx, profile.ID)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, users.ErrUserNotFound)
	}

	data := modelUserProfileForUpdateFromUpdateProfileDto(profile)

	updatedProfile, err := s.UserRepo.UpdateProfile(ctx, data)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	return updatedProfile, nil
}
