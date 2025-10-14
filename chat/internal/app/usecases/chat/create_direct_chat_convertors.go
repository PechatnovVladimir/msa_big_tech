package chat

import (
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
)

func fromCreateDirectChatIN(input *dto.CreateDirectChatIN) string {
	return input.ParticipantID
}

func toCreateDirectChatOUT(input *chat.Chat) *dto.CreateDirectChatOUT {
	return &dto.CreateDirectChatOUT{
		ChatID: input.ChatID,
	}
}
