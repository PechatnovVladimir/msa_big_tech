package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
)

func (r *RepositoryInMemory) SearchByNickname(ctx context.Context, nickname string) ([]*users.UserProfile, error) {
	for _, userProfile := range r.storage {
		if userProfile.Nickname == nickname {
			u := make([]*users.UserProfile, 1)
			u[0] = &userProfile
			return u, nil
		}
	}
	return []*users.UserProfile{}, users.ErrUserNotFound
}
