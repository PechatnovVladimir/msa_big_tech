package chat

import (
	"context"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"github.com/google/uuid"
	"strings"
	"time"
)

// CreateDirectChat - создать личный чат
func (r *Repository) CreateDirectChat(ctx context.Context, userID string, participantID string) (*chat.Chat, error) {
	chatID := uuid.New().String()
	createdAt := time.Now()

	//chatRow, chatMembersRow := fromModelCreateDirectChat(chatID, userID, participantID)

	//вставка в таблицу chats (id чата) и в таблицу chat_members (id чата и id собеседников)
	queryChatsTable := r.sb.
		Insert(chatsTable).
		Columns(chatsTableColumns...).
		Values(chatID, createdAt).
		Suffix("RETURNING " + strings.Join(chatsTableColumns, ", "))

	queryChatMembersTable := r.sb.
		Insert(chatMembersTable).
		Columns(chatMembersTableColumns...).
		Values(chatID, userID).
		Values(chatID, participantID).
		Suffix("RETURNING " + strings.Join(chatMembersTableColumns, ", "))

	pool := r.db.GetQueryEngine(ctx)

	tagChats, err := pool.Execx(ctx, queryChatsTable)
	if err != nil {
		return nil, err
	}

	tagChatMembers, err := pool.Execx(ctx, queryChatMembersTable)
	if err != nil {
		return nil, err
	}

	if tagChats.RowsAffected() != 1 && tagChatMembers.RowsAffected() != 2 {
		return nil, fmt.Errorf("failed to create direct chat")
	}

	return toModelCreateDirectChat(chatID, userID, participantID, createdAt), nil
}
