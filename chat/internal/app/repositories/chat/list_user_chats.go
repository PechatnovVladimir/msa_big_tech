package chat

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
)

func (r *Repository) ListUserChats(ctx context.Context, userID string) ([]*chat.Chat, error) {
	//select * from chat_members where user_id = ?userID
	query := r.sb.
		Select("*").
		From(chatMembersTable).
		Where(sq.Eq{chatMembersTableColumnUserID: userID})

	pool := r.db.GetQueryEngine(ctx)

	var rowsChats []ChatMembersRow
	if err := pool.Selectx(ctx, &rowsChats, query); err != nil {
		return nil, err
	}

	_ = rowsChats

	//out := make([]*chat.Chat, 0, len(rowsChats))
	//for i := range rowsChats {
	//	out = append(out, toModelFromChatMembersRow(&rowsChats[i]))
	//}

	return nil, nil
}
