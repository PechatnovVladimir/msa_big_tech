package users

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"strings"
)

func (r *Repository) UpdateProfile(ctx context.Context, p *users.UserProfileForUpdate) (*users.UserProfile, error) {
	query := r.sb.Update(usersTable).
		Where(sq.Eq{"id": p.ID})

	if p.Email != nil {
		query = query.Set("email", p.Email)
	}

	if p.Nickname != nil {
		query = query.Set("nickname", p.Nickname)
	}

	if p.Bio != nil {
		query = query.Set("bio", p.Bio)
	}

	if p.Avatar != nil {
		query = query.Set("avatar", p.Avatar)
	}

	query = query.Suffix("RETURNING " + strings.Join(usersTableColumns, ","))

	pool := r.db.GetQueryEngine(ctx)

	var outRow UserProfileRow
	if err := pool.Getx(ctx, &outRow, query); err != nil {
		return &users.UserProfile{}, err
	}

	return ToModel(&outRow), nil

}
