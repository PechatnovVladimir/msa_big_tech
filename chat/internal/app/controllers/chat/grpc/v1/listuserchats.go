package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
	"log"
)

func (s *Service) ListUserChats(ctx context.Context, request *chat.ListUserChatsRequest) (*chat.ListUserChatsResponse, error) {
	log.Println("ChatService ListUserChats called")

	return &chat.ListUserChatsResponse{}, nil
}
