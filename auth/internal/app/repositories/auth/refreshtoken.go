package auth

import (
	"context"
	"log"
)

func (r *Repository) RefreshToken(ctx context.Context, userID string, accessToken string, refreshToken string) error {
	log.Println("refreshing token in repository")
	return nil
}
