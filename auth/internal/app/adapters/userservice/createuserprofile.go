package userservice

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/usecases/auth/dto"
	"log"
)

func (c *Client) CreateUserProfile(ctx context.Context, in dto.CreateUserProfileInDTO) (out dto.CreateUserProfileOutDTO, err error) {
	logger.Println("Create User Profile called userid:", in.UserID)
	out.UserID = in.UserID
	return out, nil
}
