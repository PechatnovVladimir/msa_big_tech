package userprovider

import (
	"context"
	social2 "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
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

func (c *Client) GetUserFromContext(ctx context.Context) (*social2.User, error) {
	_ = ctx
	//пока не из контекста, а из переменной окружения
	return &social2.User{
		UserID: os.Getenv("CurrentUser"),
	}, nil
}
