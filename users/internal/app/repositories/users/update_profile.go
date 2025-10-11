package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"log"
)

func (r *Repository) UpdateProfile(ctx context.Context, p *users.UserProfile) error {
	_ = ctx
	_ = p
	log.Println("Repository UpdateProfile called")
	return nil
}
