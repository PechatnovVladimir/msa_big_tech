package chat

import (
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
)

func fromGetChatIN(in *dto.GetChatIN) string {
	return in.ChatID
}

func toGetChatOUT(in *chat.Chat) *dto.GetChatOUT {
	return &dto.GetChatOUT{
		Chat: in,
	}
}
