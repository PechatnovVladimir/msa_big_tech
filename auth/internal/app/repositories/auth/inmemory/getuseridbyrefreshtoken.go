package inmemory

import "errors"

func (r *InMemory) GetUserIDByRefreshToken(refreshToken string) (string, error) {
	r.mx.RLock()
	defer r.mx.RUnlock()

	t, exists := r.tokenStorage[refreshToken]
	if !exists {
		return "", errors.New("refresh token not found")
	}

	return t.UserID, nil
}
