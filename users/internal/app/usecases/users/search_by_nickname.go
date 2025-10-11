package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
)

func (s *Service) SearchByNickname(ctx context.Context, d dto.SearchByNicknameDTO) ([]*users.UserProfile, error) {
	nickname := d.Query
	userProfiles, err := s.repository.SearchByNickname(ctx, nickname)
	if err != nil {
		return nil, err
	}
	return userProfiles, nil
}
