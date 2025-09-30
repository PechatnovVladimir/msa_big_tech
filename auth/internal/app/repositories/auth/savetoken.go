package auth

import (
	"context"
	"log"
)

func (r *Repository) SaveToken(ctx context.Context, userID string, accessToken string, refreshToken string) error {
	log.Println("save token in repository")
	return nil
}
