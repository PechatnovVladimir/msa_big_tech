package auth

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/lib/logger"
)

// SaveUser - сохранить пользователя
func (r *Repository) SaveUser(ctx context.Context, userID string, email string, password string) error {
	logger.Info(ctx, "save user in repository")
	return nil
}
