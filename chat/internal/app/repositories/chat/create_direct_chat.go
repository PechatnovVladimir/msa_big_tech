package chat

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"strings"
)

// CreateDirectChat - создать личный чат
func (r *Repository) CreateDirectChat(ctx context.Context, chat *chat.Chat) (*chat.Chat, error) {
	chatsRow, chatMembersRow := fromModelForCreateDirectChat(chat)

	//insert в pg
	//вставка в таблицу chats (id чата) и в таблицу chat_members (id чата и id владельца чата)
	queryChatsTable := r.sb.
		Insert(chatsTable).
		Columns(chatsTableColumns...).
		Values(chatsRow.Values()...).
		Suffix("RETURNING " + strings.Join(chatsTableColumns, ", "))

	queryChatMembersTable := r.sb.
		Insert(chatMembersTable).
		Columns(chatMembersTableColumns...).
		Values(chatMembersRow.Values()...).
		Suffix("RETURNING " + strings.Join(chatMembersTableColumns, ", "))

	pool := r.db.GetQueryEngine(ctx)

	_, err := pool.Execx(ctx, queryChatsTable)
	if err != nil {
		return nil, err
	}

	var outRow ChatMembersRow
	err = pool.Getx(ctx, &outRow, queryChatMembersTable)
	if err != nil {
		return nil, err
	}

	return toModelForCreateDirectChat(&outRow), nil
}
