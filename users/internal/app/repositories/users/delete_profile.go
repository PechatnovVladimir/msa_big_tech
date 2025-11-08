package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/lib/logger"
)

func (r *Repository) DeleteProfile(ctx context.Context, profileID string) error {
	_ = ctx
	logger.Info(ctx, "Repository DeleteProfile called")
	return nil
}
