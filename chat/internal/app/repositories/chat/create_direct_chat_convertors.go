package chat

import (
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"time"
)

//func fromModelCreateDirectChat(chatID string, userID string, participantID string) (*ChatsRow, *ChatMembersRow) {
//	return nil, nil
//}

func toModelCreateDirectChat(chatID string, userID string, participantID string, createdAt time.Time) *chat.Chat {
	return &chat.Chat{
		ChatID:   chatID,
		CreateAt: createdAt,
		Members:  []string{userID, participantID},
	}
}
