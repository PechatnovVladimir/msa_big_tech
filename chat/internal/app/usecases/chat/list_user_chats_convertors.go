package chat

import (
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
)

func fromListUserChatsIN(in *dto.ListUserChatsIN) string {
	return in.UserID
}

func toListUserChatsOUT(in []*chat.Chat) *dto.ListUserChatsOUT {
	return &dto.ListUserChatsOUT{
		Chats: in,
	}
}
