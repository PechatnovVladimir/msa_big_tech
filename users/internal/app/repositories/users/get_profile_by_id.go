package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"log"
)

func (r *Repository) GetProfileByID(ctx context.Context, profileID int64) (users.UserProfile, error) {
	_ = profileID
	_ = ctx
	log.Println("Repository GetUserProfileByID called")
	return users.UserProfile{}, nil
}
