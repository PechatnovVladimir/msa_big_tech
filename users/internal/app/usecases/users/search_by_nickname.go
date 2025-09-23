package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
)

func (s *UserService) SearchByNickname(ctx context.Context, dto dto.SearchByNicknameDTO) ([]*users.UserProfile, error) {
	return nil, nil
}
