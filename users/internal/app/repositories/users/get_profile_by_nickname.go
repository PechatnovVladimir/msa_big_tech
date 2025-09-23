package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"log"
)

func (r *Repository) GetProfileByNickname(ctx context.Context, nickname string) (users.UserProfile, error) {
	_ = nickname
	_ = ctx
	log.Println("Repository GetUserProfileByNickname called")
	return users.UserProfile{}, nil
}
