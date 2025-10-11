package auth

import (
	"context"
	"log"
)

func (r *Repository) GetUserIDByRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	log.Println("GetUserIDByRefreshToken in repository")
	return "", nil
}
