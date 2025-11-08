package auth

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/lib/logger"
)

func (r *Repository) SaveToken(ctx context.Context, userID string, accessToken string, refreshToken string) error {
	logger.Info(ctx, "save token in repository")
	return nil
}
