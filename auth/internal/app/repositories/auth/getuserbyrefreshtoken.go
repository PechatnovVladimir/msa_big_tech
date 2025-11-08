package auth

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/lib/logger"
)

func (r *Repository) GetUserIDByRefreshToken(ctx context.Context, refreshToken string) (string, error) {
	logger.Info(ctx, "GetUserIDByRefreshToken in repository")
	return "", nil
}
