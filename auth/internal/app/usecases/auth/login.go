package auth

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/usecases/auth/dto"
)

func (s *Service) Login(ctx context.Context, in dto.LoginInDTO) (out dto.LoginOutDTO, err error) {
	return out, nil
}
