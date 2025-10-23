package chat

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"github.com/PechatnovVladimir/msa_big_tech/lib/pagination"
	"strings"
)

func (r *Repository) ListMessages(ctx context.Context, chatID string, p pagination.Options) ([]*chat.Message, error) {
	query := r.sb.Select("*").
		From(messagesTable).
		Where(sq.Eq{chatMembersTableColumnChatID: chatID})

	//Сортировка
	orderSQL := buildOrderBy(p.OrderBy())
	if len(orderSQL) == 0 {
		orderSQL = []string{messagesTableColumnCreatedAt, messagesTableColumnID}
	}
	query = query.OrderBy(orderSQL...)

	//пагинация
	if l := p.Limit(); l > 0 {
		query = query.Limit(uint64(l))
	}
	//select * from messages where created_at > $1
	if t := p.Cursor().Time; t != nil {
		query = query.Where(sq.Gt{messagesTableColumnCreatedAt: t})
	}

	pool := r.db.GetQueryEngine(ctx)

	var outRow []MessageRow
	if err := pool.Selectx(ctx, &outRow, query); err != nil {
		return nil, err
	}

	out := make([]*chat.Message, 0, len(outRow))
	for i := range outRow {
		out = append(out, toModelForListMessages(outRow[i]))
	}

	return out, nil
}

// Разрешённые поля сортировки → SQL
func buildOrderBy(fields []pagination.SortField) []string {
	if len(fields) == 0 {
		return nil
	}

	whitelist := map[string]bool{
		"created_at": true,
		"sender_id":  true,
	}

	var out []string
	for _, f := range fields {
		name := strings.ToLower(strings.TrimSpace(f.Name))
		if !whitelist[name] {
			continue
		}
		dir := "ASC"
		if f.Desc {
			dir = "DESC"
		}
		out = append(out, fmt.Sprintf("%s %s", name, dir))
	}
	return out
}
