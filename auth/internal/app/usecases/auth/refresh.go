package auth

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/usecases/auth/dto"
)

func (s *Service) Refresh(ctx context.Context, in dto.RefreshInDTO) (out dto.RefreshOutDTO, err error) {

	refreshToken := in.RefreshToken

	userID, err := s.AuthRepo.GetUserIDbyRefreshToken(ctx, refreshToken)
	if err != nil {
		return out, err
	}

	accessToken, _ := generateRandomToken(32)

	err = s.AuthRepo.RefreshToken(ctx, userID, accessToken, refreshToken)

	return out, nil
}
