package users

import (
	"context"
	"log"
)

func (r *InMemory) DeleteProfile(ctx context.Context, profileID string) error {
	_ = ctx
	log.Println("Repository DeleteProfile called")
	return nil
}
