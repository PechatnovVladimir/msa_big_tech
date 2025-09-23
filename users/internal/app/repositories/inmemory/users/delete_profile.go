package users

import (
	"context"
	"log"
)

func (r *RepositoryInMemory) DeleteProfile(ctx context.Context, profileID string) error {
	_ = ctx
	log.Println("Repository DeleteProfile called")
	return nil
}
