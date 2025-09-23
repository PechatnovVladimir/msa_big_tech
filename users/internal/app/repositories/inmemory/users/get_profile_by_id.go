package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
)

func (r *RepositoryInMemory) GetProfileByID(ctx context.Context, profileID string) (*users.UserProfile, error) {
	r.mx.Lock()
	defer r.mx.Unlock()

	id := users.UserID(profileID)
	userProfile, ok := r.storage[id]
	if !ok {
		return &users.UserProfile{}, users.ErrUserNotFound
	}

	return &userProfile, nil
}
