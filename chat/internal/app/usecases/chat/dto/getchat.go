package dto

import "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"

type GetChatIN struct {
	ChatID string
}

type GetChatOUT struct {
	Chat *chat.Chat
}
