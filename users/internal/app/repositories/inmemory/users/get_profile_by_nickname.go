package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
)

func (r *InMemory) GetProfileByNickname(ctx context.Context, nickname string) (*users.UserProfile, error) {
	r.mx.Lock()
	defer r.mx.Unlock()

	for _, userProfile := range r.storage {
		if userProfile.Nickname == nickname {
			return &userProfile, nil
		}
	}
	return &users.UserProfile{}, users.ErrUserNotFound
}
