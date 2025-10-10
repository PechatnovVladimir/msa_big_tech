package chat

import "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"

func toModelFromChatMembersRow(r *ChatMembersRow) *chat.Chat {
	if r == nil {
		return nil
	}
	return &chat.Chat{
		UserID: r.UserID,
		ChatID: r.ChatID,
	}
}
