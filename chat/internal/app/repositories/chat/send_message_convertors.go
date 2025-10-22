package chat

import (
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"time"
)

type MessageRow struct {
	ID        string    `db:"id"`
	ChatID    string    `db:"chat_id"`
	SenderID  string    `db:"sender_id"`
	Text      string    `db:"text"`
	CreatedAt time.Time `db:"created_at"`
}

func (r *MessageRow) Values() []any {
	return []any{r.ID, r.ChatID, r.SenderID, r.Text, r.CreatedAt}
}

func fromModelForSendMessage(m *chat.Message) (*MessageRow, *ChatMembersRow) {
	return &MessageRow{
			ID:        m.MessageID,
			ChatID:    m.ChatID,
			SenderID:  m.SenderID,
			Text:      m.Text,
			CreatedAt: m.CreatedAt,
		},
		&ChatMembersRow{
			ChatID: m.ChatID,
			UserID: m.SenderID,
		}
}

func toModelForSendMessage(m *MessageRow) *chat.Message {
	return &chat.Message{
		MessageID: m.ID,
		ChatID:    m.ChatID,
		SenderID:  m.SenderID,
		Text:      m.Text,
		CreatedAt: m.CreatedAt,
	}
}
