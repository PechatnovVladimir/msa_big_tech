package v1

import (
	"buf.build/go/protovalidate"
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Service) GetChat(ctx context.Context, request *chat.GetChatRequest) (*chat.GetChatResponse, error) {
	log.Println("ChatService GetChat called")

	//валидация по proto описанию
	err := protovalidate.Validate(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	out, err := s.ChatUseCase.GetChat(ctx, dto.GetChatIN{
		ChatID: request.ChatId,
	})
	if err != nil {
		return nil, err
	}

	chatOut := chat.Chat{
		ChatId:      out.ChatID,
		Description: "description",
	}

	return &chat.GetChatResponse{
		Chat: &chatOut,
	}, nil
}
