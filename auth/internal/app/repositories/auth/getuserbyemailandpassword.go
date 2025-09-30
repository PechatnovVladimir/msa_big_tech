package auth

import (
	"context"
	"log"
)

func (r *Repository) GetUserByEmailAndPassword(ctx context.Context, email string, password string) (string, error) {
	log.Println("GetUserByEmailAndPassword in repository")
	return "", nil
}
