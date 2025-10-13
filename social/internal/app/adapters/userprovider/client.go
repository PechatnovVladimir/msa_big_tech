package userprovider

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social"
	"os"
)

type Client struct {
	// http/grpc/...
}

var (
	_ social.UserProvider = (*Client)(nil)
)

func New() *Client {
	return &Client{}
}

func (c *Client) GetUserFromContext(ctx context.Context) (*models.User, error) {
	_ = ctx
	//пока не из контекста, а из переменной окружения
	return &models.User{
		UserID: os.Getenv("CurrentUser"),
	}, nil
}
