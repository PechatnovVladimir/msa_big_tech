package server

import (
	"context"
	chat "github.com/PechatnovVladimir/msa_big_tech/chat/pkg/api"
	"google.golang.org/grpc"
	"log"
)

type ChatService struct {
	chat.UnimplementedChatServiceServer
}

func NewChatService() *ChatService {
	return &ChatService{}
}

func (s *ChatService) CreateDirectChat(ctx context.Context, request *chat.CreateDirectChatRequest) (*chat.CreateDirectChatResponse, error) {
	return &chat.CreateDirectChatResponse{}, nil
}

func (s *ChatService) GetChat(ctx context.Context, request *chat.GetChatRequest) (*chat.GetChatResponse, error) {
	c := &chat.Chat{
		ChatId: request.ChatId,
		Name:   "chat....",
	}
	return &chat.GetChatResponse{Chat: c}, nil
}

func (s *ChatService) ListUserChats(ctx context.Context, request *chat.ListUserChatsRequest) (*chat.ListUserChatsResponse, error) {
	return &chat.ListUserChatsResponse{}, nil
}
func (s *ChatService) ListChatMembers(ctx context.Context, request *chat.ListChatMembersRequest) (*chat.ListChatMembersResponse, error) {
	log.Println("ListChatMembers called")
	a := []int64{1, 2, 3, 4, 5}
	return &chat.ListChatMembersResponse{UserIds: a}, nil
	//return &chat.ListChatMembersResponse{}, nil
}

func (s *ChatService) SendMessage(ctx context.Context, request *chat.SendMessageRequest) (*chat.SendMessageResponse, error) {
	return &chat.SendMessageResponse{}, nil
}

func (s *ChatService) ListMessages(ctx context.Context, request *chat.ListMessagesRequest) (*chat.ListMessagesResponse, error) {
	return &chat.ListMessagesResponse{}, nil
}

func (s *ChatService) StreamMessagesStreamMessages(*chat.StreamMessagesRequest, grpc.ServerStreamingServer[chat.StreamMessagesResponse]) error {
	return nil
}
