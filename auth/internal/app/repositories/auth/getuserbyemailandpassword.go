package auth

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/lib/logger"
)

func (r *Repository) GetUserByEmailAndPassword(ctx context.Context, email string, password string) (string, error) {
	logger.Info(ctx, "GetUserByEmailAndPassword in repository")
	return "", nil
}
