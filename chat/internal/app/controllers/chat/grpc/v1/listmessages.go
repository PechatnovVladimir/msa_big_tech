package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
	"log"
)

func (s *Service) ListMessages(ctx context.Context, request *chat.ListMessagesRequest) (*chat.ListMessagesResponse, error) {
	log.Println("ChatService ListMessages called")
	return &chat.ListMessagesResponse{}, nil
}
