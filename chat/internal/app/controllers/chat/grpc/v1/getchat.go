package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
	"log"
)

func (s *Service) GetChat(ctx context.Context, request *chat.GetChatRequest) (*chat.GetChatResponse, error) {
	log.Println("ChatService GetChat called")
	c := &chat.Chat{
		ChatId: request.ChatId,
	}
	return &chat.GetChatResponse{Chat: c}, nil
}
