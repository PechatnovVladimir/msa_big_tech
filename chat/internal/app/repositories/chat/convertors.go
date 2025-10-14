package chat

import "time"

type ChatsRow struct {
	ID        string    `db:"id"`
	CreatedAt time.Time `db:"created_at"`
}

type ChatMembersRow struct {
	ChatID string `db:"chat_id"`
	UserID string `db:"user_id"`
}

type ChatsMembersRow struct {
	ChatID    string    `db:"chat_id"`
	CreatedAt time.Time `db:"created_at"`
	UserIDs   []string  `db:"user_ids"`
}

func (r *ChatsRow) Values() []any {
	return []any{r.ID, r.CreatedAt}
}

func (r *ChatMembersRow) Values() []any {
	return []any{r.ChatID, r.UserID}
}
