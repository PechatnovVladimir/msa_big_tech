package auth

import "context"

func (r *Repository) GetUserIDByRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	return "", nil
}
