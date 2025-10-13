package social

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
)

func (r *Repository) CreateFriend(ctx context.Context, fr *social.FriendRequest) error {
	friendRow := fromModelForCreateFriend(fr)
	query := r.sb.
		Insert(friendsTable).
		Columns(friendsTableColumns...).
		Values(friendRow.Values()...)

	pool := r.db.GetQueryEngine(ctx)

	_, err := pool.Execx(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
