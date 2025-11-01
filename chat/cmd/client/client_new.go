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
	return &ClientNew{
		conn: client,
	}, nil

}

func (c *ClientNew) Close() error {
	c.conn.
}
