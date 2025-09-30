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

// CreateDirectChat - delivery создать личный чат
func (s *Service) CreateDirectChat(ctx context.Context, request *chat.CreateDirectChatRequest) (*chat.CreateDirectChatResponse, error) {
	log.Println("grps ChatService CreateDirectChat called")

	//валидация по proto описанию
	err := protovalidate.Validate(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	inUC := grpc2dto(request)

	//вызов бизнес-логики
	outUC, err := s.ChatUseCase.CreateDirectChat(ctx, inUC)
	if err != nil {
		return nil, err
	}

	out := dto2grpc(outUC)

	return out, nil
}

// конвертация request в dto
func grpc2dto(request *chat.CreateDirectChatRequest) dto.CreateDirectChatIN {
	out := dto.CreateDirectChatIN{}
	out.ParticipantID = request.ParticipantId
	return out
}

// конвертация dto в response
func dto2grpc(in dto.CreateDirectChatOUT) *chat.CreateDirectChatResponse {
	out := chat.CreateDirectChatResponse{
		ChatId: in.ChatID,
	}
	return &out
}
