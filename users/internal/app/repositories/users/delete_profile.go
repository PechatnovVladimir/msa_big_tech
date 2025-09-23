package users

import (
	"context"
	"log"
)

func (r *Repository) DeleteProfile(ctx context.Context, profileID int64) error {
	_ = ctx
	log.Println("Repository DeleteProfile called")
	return nil
}
