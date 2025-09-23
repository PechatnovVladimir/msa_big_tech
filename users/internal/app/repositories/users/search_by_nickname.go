package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"log"
)

func (r *Repository) SearchByNickname(ctx context.Context, nickname string) ([]*users.UserProfile, error) {
	_ = ctx
	_ = nickname
	log.Println("Repository SearchByNickname called")
	return []*users.UserProfile{}, nil
}
