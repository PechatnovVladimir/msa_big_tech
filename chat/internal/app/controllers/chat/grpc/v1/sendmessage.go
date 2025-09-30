package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
	"log"
)

func (s *Service) SendMessage(ctx context.Context, request *chat.SendMessageRequest) (*chat.SendMessageResponse, error) {
	log.Println("ChatService SendMessage called")
	return &chat.SendMessageResponse{}, nil
}
