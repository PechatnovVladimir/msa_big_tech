package chat

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetChat(ctx context.Context, chatID string) (*chat.Chat, error) {
	query := r.sb.
		Select("chat_id,created_at,coalesce(array_agg(user_id),'{}') user_ids").
		From("chat_members").
		Join("chats on chats.id=chat_members.chat_id").
		Where(sq.Eq{chatsTableColumnID: chatID}).
		GroupBy("chat_id, created_at")

	pool := r.db.GetQueryEngine(ctx)

	var outRow ChatsMembersRow
	if err := pool.Getx(ctx, &outRow, query); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, chat.ErrNotFound
		}
		return nil, err
	}

	return toModelGetChat(outRow), nil

}
