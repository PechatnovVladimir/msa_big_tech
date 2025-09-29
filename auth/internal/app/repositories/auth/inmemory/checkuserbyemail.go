package inmemory

import (
	"context"
	"errors"
)

func (r *InMemory) CheckUserByEmail(ctx context.Context, email string) (err error) {

	r.mx.RLock()
	defer r.mx.RUnlock()

	_, exists := r.userStorage[email]
	if exists {
		return errors.New("user already exists")
	}

	return nil
}
