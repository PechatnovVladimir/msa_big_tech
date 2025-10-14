package chat

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"github.com/google/uuid"
	"strings"
)

// CreateDirectChat - создать личный чат
func (r *Repository) CreateDirectChat(ctx context.Context, userID string, participantID string) (*chat.Chat, error) {
	chatID := uuid.New().String()

	//chatRow, chatMembersRow := fromModelForCreateDirectChat(chatID, userID, participantID)

	//insert в pg
	//вставка в таблицу chats (id чата) и в таблицу chat_members (id чата и id собеседников)
	queryChatsTable := r.sb.
		Insert(chatsTable).
		Columns(chatsTableColumns...).
		Values(chatID).
		Suffix("RETURNING " + strings.Join(chatsTableColumns, ", "))

	queryChatMembersTable := r.sb.
		Insert(chatMembersTable).
		Columns(chatMembersTableColumns...).
		Values(chatID, userID).
		Values(chatID, participantID).
		Suffix("RETURNING " + strings.Join(chatMembersTableColumns, ", "))

	pool := r.db.GetQueryEngine(ctx)

	var outChatsRow ChatsRow
	err := pool.Getx(ctx, &outChatsRow, queryChatsTable)
	if err != nil {
		return nil, err
	}

	var outChatMembersRow []ChatMembersRow
	err = pool.Getx(ctx, &outChatMembersRow, queryChatMembersTable)
	if err != nil {
		return nil, err
	}

	return toModelFromChatMembersRow(&outChatsRow, &outChatMembersRow), nil
}
