package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"strings"
)

func (r *Repository) CreateProfile(ctx context.Context, p *users.UserProfile) error {
	row := FromModel(p)

	query := r.sb.
		Insert(usersTable).
		Columns(usersTableColumns...).
		Values(row.Values()...).
		Suffix("RETURNING " + strings.Join(usersTableColumns, ","))

	pool := r.db.GetQueryEngine(ctx)

	var outRow UserProfileRow
	if err := pool.Getx(ctx, &outRow, query); err != nil {
		return err
	}

	return nil
}
