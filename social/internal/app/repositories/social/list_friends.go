package social

import (
	"context"
	"fmt"
	sq "github.com/Masterminds/squirrel"
	"github.com/PechatnovVladimir/msa_big_tech/pkg/pagination"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
	"log"
	"strings"
)

func (r *Repository) ListFriends(ctx context.Context, userID string, p pagination.Options) ([]*social.Friend, error) {

	log.Println("Repo - ", userID, p.Limit(), p.Cursor())

	query := r.sb.
		Select("user_id,friend_user_id,created_at").
		From("friends").
		Where(sq.Eq{"user_id": userID})

	//Сортировка
	orderSQL := buildOrderBy(p.OrderBy())
	if len(orderSQL) == 0 {
		orderSQL = []string{"created_at", "friend_user_id"}
	}
	query = query.OrderBy(orderSQL...)

	//пагинация
	if l := p.Limit(); l > 0 {
		query = query.Limit(uint64(l))
	}
	//select ... from friends where created_at > $1
	if t := p.Cursor().Time; t != nil {
		query = query.Where(sq.Gt{"created_at": t})
	}

	pool := r.db.GetQueryEngine(ctx)

	log.Println(query.ToSql())

	var outRow []FriendRow
	if err := pool.Selectx(ctx, &outRow, query); err != nil {
		return nil, err
	}

	out := make([]*social.Friend, 0, len(outRow))
	for i := range outRow {
		out = append(out, toModelForListFriends(outRow[i]))
	}

	return out, nil
}

func buildOrderBy(fields []pagination.SortField) []string {
	if len(fields) == 0 {
		return nil
	}

	whitelist := map[string]bool{
		"created_at":     true,
		"friend_user_id": true,
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
