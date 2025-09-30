package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
	"log"
)

func (s *Service) ListChatMembers(ctx context.Context, request *chat.ListChatMembersRequest) (*chat.ListChatMembersResponse, error) {
	log.Println("ChatService ListChatMembers called")
	return &chat.ListChatMembersResponse{UserIds: []string{"CEDD5E54-997C-42D2-9AD5-002F001BA300"}}, nil
}
