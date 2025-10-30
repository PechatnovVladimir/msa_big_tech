package chat

import (
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
)

func fromSendMessageIN(in *dto.SendMessageIN) *chat.Message {
	return &chat.Message{
		ChatID: in.ChatID,
		Text:   in.Text,
	}
}

func toSendMessageOUT(in *chat.Message) *dto.SendMessageOUT {
	return &dto.SendMessageOUT{
		MessageID: in.MessageID,
		ChatID:    in.ChatID,
		UserID:    in.SenderID,
		CreatedAt: in.CreatedAt,
		Text:      in.Text,
	}
}
