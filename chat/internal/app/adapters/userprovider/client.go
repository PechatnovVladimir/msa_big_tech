package userprovider

import (
	models "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat"
	"golang.org/x/net/context"
	"os"
)

type Client struct {
	// http/grpc/...
}

var (
	_ chat.UserProvider = (*Client)(nil)
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
