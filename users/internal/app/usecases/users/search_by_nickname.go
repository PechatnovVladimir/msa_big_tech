package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
)

func (s *Service) SearchByNickname(ctx context.Context, in dto.SearchByNicknameDTO) ([]*users.UserProfile, error) {

	query := users.UserProfileFilter{}

	if in.Query.IDs != nil {
		query.IDs = in.Query.IDs
	}

	if in.Query.Nickname != nil {
		query.Nickname = in.Query.Nickname
	}

	if in.Query.Email != nil {
		query.Email = in.Query.Email
	}

	if in.Query.CreatedFrom != nil {
		query.CreatedFrom = in.Query.CreatedFrom
	}

	if in.Query.CreatedTo != nil {
		query.CreatedTo = in.Query.CreatedTo
	}

	limit := in.Limit

	userProfiles, err := s.repository.SearchByNickname(ctx, &query, &limit)

	if err != nil {
		return nil, err
	}
	return userProfiles, nil
}
