package userservice

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/usecases/auth/dto"
	"github.com/PechatnovVladimir/msa_big_tech/lib/logger"
)

func (c *Client) CreateUserProfile(ctx context.Context, in dto.CreateUserProfileInDTO) (out dto.CreateUserProfileOutDTO, err error) {
	logger.Info(ctx, "Create User Profile called userid:", in.UserID)
	out.UserID = in.UserID
	return out, nil
}
