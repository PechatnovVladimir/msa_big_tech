package userservice

import (
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social"
	"log"
)

type Client struct {
	// http/grpc/...
}

var (
	_ social.UserService = (*Client)(nil)
)

func New() *Client {
	return &Client{}
}

func (c *Client) Test() error {
	log.Println("user service Test() called")
	return nil
}
