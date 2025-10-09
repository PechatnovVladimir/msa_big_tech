package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
)

func (s *Service) SearchByNickname(ctx context.Context, in dto.SearchByNickname) ([]*users.UserProfile, error) {

	query, limit := getUserProfileFilterFromSearchByNickNameDto(in)

	userProfiles, err := s.repository.SearchByNickname(ctx, &query, &limit)

	if err != nil {
		return nil, users.ErrUserNotFound
	}

	return userProfiles, nil
}
