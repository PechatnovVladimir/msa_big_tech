package chat

import "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"

func toModelForListMessages(row MessageRow) *chat.Message {
	return &chat.Message{
		MessageID: row.ID,
		Text:      row.Text,
		CreatedAt: row.CreatedAt,
		ChatID:    row.ChatID,
		SenderID:  row.SenderID,
	}
}
