package chat

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"github.com/jackc/pgx/v5"
	"log"
)

func (r *Repository) GetChat(ctx context.Context, chatID string) (*chat.Chat, error) {
	query := r.sb.
		Select("chat_members.chat_id", "chat_members.user_id", "chats.created_at").
		From("chats").
		Join("chat_members on chat_members.chat_id = chats.id").
		Where(sq.Eq{chatsTableColumnID: chatID})

	pool := r.db.GetQueryEngine(ctx)

	log.Println(query.ToSql())

	var outRow []ChatsMembersRow
	if err := pool.Selectx(ctx, &outRow, query); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, chat.ErrNotFound
		}
		return nil, err
	}

	_ = outRow

	return toModelGetChat(outRow), nil

}
