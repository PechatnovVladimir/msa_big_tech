package users

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"strings"
)

func (r *Repository) SearchByNickname(ctx context.Context, f *users.UserProfileFilter, limit *uint64) ([]*users.UserProfile, error) {
	if f == nil {
		f = &users.UserProfileFilter{}
	}

	q := r.sb.
		Select(usersTableColumns...).
		From(usersTable)

	//включаем фильтры
	if len(f.IDs) > 0 {
		q = q.Where(sq.Eq{"id": f.IDs})
	}

	if v := f.Email; v != nil && strings.TrimSpace(*v) != "" {
		q = q.Where("lower(email) = lower(?)", strings.TrimSpace(*v))
	}

	if v := f.Nickname; v != nil && strings.TrimSpace(*v) != "" {
		q = q.Where("nickname ILIKE ?", "%"+strings.TrimSpace(*v)+"%")
	}

	if v := f.CreatedFrom; v != nil {
		q = q.Where(sq.GtOrEq{"created_at": *v})
	}

	if v := f.CreatedTo; v != nil {
		q = q.Where(sq.Lt{"created_at": *v})
	}

	if limit != nil {
		q = q.Limit(*limit)
	}

	pool := r.db.GetQueryEngine(ctx)

	var rows []UserProfileRow
	if err := pool.Selectx(ctx, &rows, q); err != nil {
		return nil, err
	}

	out := make([]*users.UserProfile, 0, len(rows))
	for i := range rows {
		out = append(out, ToModel(&rows[i]))
	}
	return out, nil

}
