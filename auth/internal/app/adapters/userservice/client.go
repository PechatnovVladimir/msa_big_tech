package userservice

import "github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/usecases/auth"

type Client struct {
	// http/grpc/...
}

var (
	_ auth.UserService = (*Client)(nil)
)

func NewClient( /**/ ) *Client {
	return &Client{}
}
