package chat

import (
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"github.com/google/uuid"
)

func fromCreateDirectChatIN(input *dto.CreateDirectChatIN) *chat.Chat {
	return &chat.Chat{
		UserID: input.ParticipantID,
		ChatID: uuid.New().String(),
	}
}

func toCreateDirectChatOUT(input *chat.Chat) *dto.CreateDirectChatOUT {
	return &dto.CreateDirectChatOUT{
		ChatID: input.ChatID,
	}
}
