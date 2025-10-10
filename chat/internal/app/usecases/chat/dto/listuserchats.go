package dto

import "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"

type ListUserChatsIN struct {
	UserID string
}
type ListUserChatsOUT struct {
	Chats []*chat.Chat
}
