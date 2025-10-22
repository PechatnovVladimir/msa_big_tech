package users

import (
	"context"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
)

func (s *Service) SearchByNickname(ctx context.Context, in dto.SearchByNickname) ([]*users.UserProfile, error) {
	const api = "UserService.SearchByNickname"

	query, limit := getUserProfileFilterFromSearchByNickNameDto(in)

	const searchDefaultLimit = 10
	if limit == 0 {
		limit = searchDefaultLimit
	}

	userProfiles, err := s.UserRepo.SearchByNickname(ctx, &query, &limit)

	if err != nil {
		return nil, fmt.Errorf("%s: %w", api, err)
	}

	if len(userProfiles) == 0 {
		return nil, fmt.Errorf("%s: %w", api, users.ErrUserNotFound)
	}

	return userProfiles, nil
}
