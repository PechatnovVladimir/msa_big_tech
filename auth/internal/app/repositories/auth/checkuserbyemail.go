package auth

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/lib/logger"
)

// CheckUserByEmail - проверить существование пользователя по e-mail
func (r *Repository) CheckUserByEmail(ctx context.Context, email string) error {
	logger.Info(ctx, "CheckUserByEmail in repository")
	return nil
}
