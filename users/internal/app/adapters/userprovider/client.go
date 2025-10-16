package userprovider

import (
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users"
	"golang.org/x/net/context"
	"os"
)

type Client struct {
	// http/grpc/...
}

var (
	_ users.UserProvider = (*Client)(nil)
)

func New() *Client {
	return &Client{}
}

func (c *Client) GetUserFromContext(ctx context.Context) (string, error) {
	_ = ctx
	//пока не из контекста, а из переменной окружения
	return os.Getenv("CurrentUser"), nil
}
