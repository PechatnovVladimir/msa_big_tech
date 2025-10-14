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
		Select("*").
		From(chatsTable).
		Where(sq.Eq{chatsTableColumnID: chatID})

	pool := r.db.GetQueryEngine(ctx)

	log.Println(query.ToSql())

	var outRow ChatsRow
	if err := pool.Getx(ctx, &outRow, query); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, chat.ErrNotFound
		}
		return nil, err
	}

	return toModelFromChatsRow(&outRow), nil

}
