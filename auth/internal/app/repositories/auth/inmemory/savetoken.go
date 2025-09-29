package inmemory

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/models/auth"
)

func (s *InMemory) SaveToken(ctx context.Context, userID string, accessToken string, refreshToken string) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	t := auth.Token{
		UserID:       userID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	s.tokenStorage[refreshToken] = t
	return nil
}
