package v1

import (
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
)

func (s *Service) StreamMessages(req *chat.StreamMessagesRequest, stream chat.ChatService_StreamMessagesServer) error {
	messageChanel, err := s.ChatUseCase.StreamMessages(
		stream.Context(),
		&dto.StreamMessagesIN{
			ChatID:           req.GetChatId(),
			SinceMessageTime: req.GetSinceMessageTime().AsTime().UTC(),
		})
	if err != nil {
		return err
	}

	for message := range messageChanel {
		if err = stream.Send(&chat.StreamMessagesResponse{Message: modelMessageToChatPB(message)}); err != nil {
			return err
		}
	}

	return nil
}
