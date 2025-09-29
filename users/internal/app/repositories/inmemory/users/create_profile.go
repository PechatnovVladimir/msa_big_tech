package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
)

func (r *InMemory) CreateProfile(ctx context.Context, p *users.UserProfile) error {
	r.mx.Lock()
	defer r.mx.Unlock()

	r.storage[p.ID] = users.UserProfile{
		ID:       p.ID,
		Nickname: p.Nickname,
		Avatar:   p.Avatar,
		Bio:      p.Bio,
	}
	return nil
}
