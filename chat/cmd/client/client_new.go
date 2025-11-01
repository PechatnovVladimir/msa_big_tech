package main

import (
	chatPB "github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
	"github.com/PechatnovVladimir/msa_big_tech/lib/config"
	"github.com/PechatnovVladimir/msa_big_tech/lib/grpc"
)

type ClientNew struct {
	client chatPB.ChatServiceClient
	conn   *grpc.Client
}

func NewClientNew(target string, cfg *config.Config) (*ClientNew, error) {
	client, err := grpc.NewGRPCClientConn(target, cfg)
	if err != nil {
		return nil, err
	}

	c := chatPB.NewChatServiceClient(client.GrpcClient)

	return &ClientNew{
		client: c,
		conn:   client,
	}, nil

}

func (c *ClientNew) Close() error {
	_ = c.Close()
	return nil
}
