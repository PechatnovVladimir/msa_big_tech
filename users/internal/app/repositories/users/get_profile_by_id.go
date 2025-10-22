package users

import (
	"context"
	"errors"
	sq "github.com/Masterminds/squirrel"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"github.com/jackc/pgx/v5"
)

func (r *Repository) GetProfileByID(ctx context.Context, profileID string) (*users.UserProfile, error) {
	query := r.sb.
		Select("*").
		From(usersTable).
		Where(sq.Eq{"id": profileID})

	pool := r.db.GetQueryEngine(ctx)

	var outRow UserProfileRow
	if err := pool.Getx(ctx, &outRow, query); err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, users.ErrUserNotFound
		}
		return nil, err
	}

	return ToModel(&outRow), nil
}
