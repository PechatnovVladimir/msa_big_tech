package inmemory

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/models/auth"
)

func (s *InMemory) RefreshToken(ctx context.Context, userID string, accessToken string, refreshToken string) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	delete(s.tokenStorage, refreshToken)

	t := auth.Token{
		UserID:       userID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	s.tokenStorage[refreshToken] = t
	return nil
}
