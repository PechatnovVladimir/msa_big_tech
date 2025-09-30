package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
	"log"
)

func (s *Service) CreateDirectChat(ctx context.Context, request *chat.CreateDirectChatRequest) (*chat.CreateDirectChatResponse, error) {
	log.Println("ChatService CreateDirectChat called")
	return &chat.CreateDirectChatResponse{}, nil
}
