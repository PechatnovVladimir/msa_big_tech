package social

import (
	"context"
	"log"
)

func (r *Repository) RemoveFriend(ctx context.Context, userID string, friendID string) error {
	query := r.sb.Delete(friendsTable).
		Where("user_id = ?", userID).
		Where("friend_user_id = ?", friendID)

	log.Println(query.ToSql())

	pool := r.db.GetQueryEngine(ctx)

	_, err := pool.Execx(ctx, query)
	if err != nil {
		return err
	}

	return nil

}
