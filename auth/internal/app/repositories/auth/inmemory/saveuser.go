package inmemory

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/models/auth"
)

func (s *InMemory) SaveUser(ctx context.Context, userID string, email string, password string) error {
	s.mx.Lock()
	defer s.mx.Unlock()

	u := auth.User{
		UserID:   userID,
		Email:    email,
		Password: password,
	}

	s.userStorage[email] = u
	return nil
}
