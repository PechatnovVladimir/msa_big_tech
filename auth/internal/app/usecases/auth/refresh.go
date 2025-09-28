package auth

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/usecases/auth/dto"
)

func (s *Service) Refresh(ctx context.Context, in dto.RefreshInDTO) (out dto.RefreshOutDTO, err error) {
	return out, nil
}
