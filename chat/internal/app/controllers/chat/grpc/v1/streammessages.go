package v1

import (
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
	"google.golang.org/grpc"
)

func (s *Service) StreamMessages(*chat.StreamMessagesRequest, grpc.ServerStreamingServer[chat.StreamMessagesResponse]) error {

	//тестовый вызов usecase
	_, _ = s.ChatUseCase.StreamMessages(nil, dto.StreamMessagesIN{})

	return nil
}
