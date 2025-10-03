package authservice

import (
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social"
	"github.com/google/uuid"
	"log"
)

type Client struct {
	// http/grpc/...
}

var (
	_ social.AuthService = (*Client)(nil)
)

func New() *Client {
	return &Client{}
}

func (c *Client) GetAuthUser() (string, error) {
	log.Println("user service GetAuthUser() called")
	//возвращаем случайный ID
	return uuid.New().String(), nil
}
