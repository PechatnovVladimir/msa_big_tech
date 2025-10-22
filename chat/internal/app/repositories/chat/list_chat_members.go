package chat

import (
	"context"
	sq "github.com/Masterminds/squirrel"
)

func (r *Repository) ListChatMembers(ctx context.Context, chatID string) ([]*string, error) {
	//select user_id from chat_members where chat_id = ?chatID
	query := r.sb.
		Select("*").
		From(chatMembersTable).
		Where(sq.Eq{chatMembersTableColumnChatID: chatID})

	pool := r.db.GetQueryEngine(ctx)

	var rowsChats []ChatMembersRow
	if err := pool.Selectx(ctx, &rowsChats, query); err != nil {
		return nil, err
	}

	out := make([]*string, 0, len(rowsChats))
	for i := range rowsChats {
		out = append(out, &rowsChats[i].UserID)
	}
	return out, nil

}
