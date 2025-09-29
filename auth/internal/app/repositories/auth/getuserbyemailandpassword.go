package auth

import "context"

func (r *Repository) GetUserByEmailAndPassword(ctx context.Context, email string, password string) (string, error) {
	return "", nil
}
