package social

import (
	"context"
)

func (r *Repository) RemoveFriend(ctx context.Context, userID string, friendID string) error {
	query := r.sb.Delete(friendsTable).
		Where("user_id = ?", userID).
		Where("friend_user_id = ?", friendID)

	pool := r.db.GetQueryEngine(ctx)

	_, err := pool.Execx(ctx, query)
	if err != nil {
		return err
	}

	return nil

}
