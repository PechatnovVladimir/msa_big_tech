package auth

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/lib/logger"
)

func (r *Repository) RefreshToken(ctx context.Context, userID string, accessToken string, refreshToken string) error {
	logger.Info(ctx, "refreshing token in repository")
	return nil
}
