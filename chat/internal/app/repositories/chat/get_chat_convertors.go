package chat

import "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"

func toModelGetChat(r ChatsMembersRow) *chat.Chat {

	return &chat.Chat{
		ChatID:   r.ChatID,
		CreateAt: r.CreatedAt,
		Members:  r.UserIDs,
	}
}
