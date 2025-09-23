package userservice

import (
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat"
	"log"
)

type Client struct {
	// http/grpc/...
}

var (
	_ chat.UserService = (*Client)(nil)
)

func New() *Client {
	return &Client{}
}

func (c *Client) Test() error {
	log.Println("user service Test() called")
	return nil
}
