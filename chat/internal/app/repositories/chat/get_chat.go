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
		Select("*").
		From(chatMembersTable).
		Where(sq.Eq{chatMembersTableColumnChatID: chatID})

	pool := r.db.GetQueryEngine(ctx)

	var outRow ChatMembersRow
	if err := pool.Getx(ctx, &outRow, query); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, chat.ErrChatNotFound
		}
		return nil, err
	}

	return toModelFromChatMembersRow(&outRow), nil

}
