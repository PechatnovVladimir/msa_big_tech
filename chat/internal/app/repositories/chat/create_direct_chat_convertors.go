package chat

import "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"

type ChatsRow struct {
	ID string `db:"id"`
}

type ChatMembersRow struct {
	ChatID string `db:"chat_id"`
	UserID string `db:"user_id"`
}

func (r *ChatsRow) Values() []any {
	return []any{r.ID}
}

func (r *ChatMembersRow) Values() []any {
	return []any{r.ChatID, r.UserID}
}

func fromModelForCreateDirectChat(r *chat.Chat) (*ChatsRow, *ChatMembersRow) {
	if r == nil {
		return nil, nil
	}

	chatsRow := &ChatsRow{
		ID: r.ChatID,
	}

	chatMembersRow := &ChatMembersRow{
		ChatID: r.ChatID,
		UserID: r.UserID,
	}

	return chatsRow, chatMembersRow
}
