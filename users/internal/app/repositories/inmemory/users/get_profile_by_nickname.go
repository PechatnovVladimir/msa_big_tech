package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
)

func (r *RepositoryInMemory) GetProfileByNickname(ctx context.Context, nickname string) (*users.UserProfile, error) {
	for _, userProfile := range r.storage {
		if userProfile.Nickname == nickname {
			return &userProfile, nil
		}
	}
	return &users.UserProfile{}, users.ErrUserNotFound
}
