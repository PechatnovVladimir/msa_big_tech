package auth

import (
	"context"
	"log"
)

// SaveUser - сохранить пользователя
func (r *Repository) SaveUser(ctx context.Context, userID string, email string, password string) error {
	log.Println("save user in repository")
	return nil
}
