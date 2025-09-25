package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
)

func (r *InMemory) SearchByNickname(ctx context.Context, nickname string) ([]*users.UserProfile, error) {
	r.mx.Lock()
	defer r.mx.Unlock()

	for _, userProfile := range r.storage {
		if userProfile.Nickname == nickname {
			u := make([]*users.UserProfile, 1)
			u[0] = &userProfile
			return u, nil
		}
	}
	return []*users.UserProfile{}, users.ErrUserNotFound
}
