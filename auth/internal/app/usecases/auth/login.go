package auth

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/models/auth"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/usecases/auth/dto"
)

func (s *Service) Login(ctx context.Context, in dto.LoginInDTO) (out dto.LoginOutDTO, err error) {

	email := in.Email
	password := cachePassword(in.Password)

	userid, err := s.AuthRepo.GetUserByEmailAndPassword(ctx, email, password)
	if err != nil {
		return dto.LoginOutDTO{}, auth.ErrAuthUnauthenticated
	}

	out.UserID = userid
	out.AccessToken, _ = generateRandomToken(32)
	out.RefreshToken, _ = generateRandomToken(32)

	return out, nil

}
