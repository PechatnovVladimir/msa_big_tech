package main

import (
	"context"
	chatPB "github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type Client struct {
	client chatPB.ChatServiceClient
	conn   *grpc.ClientConn
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.NewClient(url,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithUnaryInterceptor(timeoutInterceptor))
	if err != nil {
		return nil, err
	}
	client := chatPB.NewChatServiceClient(conn)

	return &Client{client: client, conn: conn}, nil

}

func (c *Client) Close() error {
	return c.conn.Close()
}

func timeoutInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn,
	invoker grpc.UnaryInvoker, opts ...grpc.CallOption,
) error {
	// Устанавливаем таймаут для каждого вызова
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Вызываем метод
	return invoker(ctx, method, req, reply, cc, opts...)
}
