package v1

import (
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
)

func fromSendMessageRequestToDto(m *chat.SendMessageRequest) *dto.SendMessageIN {
	return &dto.SendMessageIN{
		ChatID: m.ChatId,
		Text:   m.Text,
	}
}

func fromDtoToSendMessageResponse(m *dto.SendMessageOUT) *chat.SendMessageResponse {
	message := &chat.Message{
		MessageId: m.MessageID,
		Text:      m.Text,
	}

	return &chat.SendMessageResponse{
		Message: message,
	}
}
