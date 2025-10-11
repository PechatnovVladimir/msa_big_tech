package auth

import (
	"context"
	"log"
)

// CheckUserByEmail - проверить существование пользователя по e-mail
func (r *Repository) CheckUserByEmail(ctx context.Context, email string) error {
	log.Println("CheckUserByEmail in repository")
	return nil
}
