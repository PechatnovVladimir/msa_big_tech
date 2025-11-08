package authprovider

import (
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social"
	"github.com/google/uuid"
)

type Client struct {
	// http/grpc/...
}

var (
	_ social.AuthProvider = (*Client)(nil)
)

func New() *Client {
	return &Client{}
}

func (c *Client) GetAuthUser() (string, error) {
	//возвращаем случайный ID
	return uuid.New().String(), nil
}
