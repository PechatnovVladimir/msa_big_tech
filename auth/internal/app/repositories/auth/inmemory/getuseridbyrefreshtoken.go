package inmemory

import (
	"context"
	"errors"
)

func (r *InMemory) GetUserIDByRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	r.mx.RLock()
	defer r.mx.RUnlock()

	t, exists := r.tokenStorage[refreshToken]
	if !exists {
		return "", errors.New("refresh token not found")
	}

	return t.UserID, nil
}
