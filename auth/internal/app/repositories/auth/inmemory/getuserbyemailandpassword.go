package inmemory

import (
	"context"
	"errors"
)

func (r *InMemory) GetUserByEmailAndPassword(ctx context.Context, email string, password string) (string, error) {
	r.mx.RLock()
	defer r.mx.RUnlock()

	u, exists := r.userStorage[email]
	if !exists || u.Password != password {
		return "", errors.New("user not authenticated")
	}

	return u.UserID, nil
}
