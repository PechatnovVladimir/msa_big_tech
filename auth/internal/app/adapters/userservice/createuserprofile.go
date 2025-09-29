package userservice

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/usecases/auth/dto"
)

func (c *Client) CreateUserProfile(ctx context.Context, in dto.CreateUserProfileInDTO) (out dto.CreateUserProfileOutDTO, err error) {
	return dto.CreateUserProfileOutDTO{}, nil
}
