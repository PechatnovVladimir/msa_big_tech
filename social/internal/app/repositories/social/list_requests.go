package social

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
	"log"
)

func (r *Repository) ListRequests(ctx context.Context, userID string) ([]*social.FriendRequest, error) {
	query := r.sb.Select("*").
		From(friendRequestsTable).
		Where(sq.Eq{friendRequestsTableColumnFromUserID: userID}).
		OrderBy("created_at DESC")

	pool := r.db.GetQueryEngine(ctx)

	var rowsFriendRequests []FriendRequestsRow

	log.Println(query.ToSql())

	err := pool.Selectx(ctx, &rowsFriendRequests, query)

	if err != nil {
		return nil, err
	}

	out := make([]*social.FriendRequest, 0, len(rowsFriendRequests))
	for i := range rowsFriendRequests {
		out = append(out, toModelForSendFriendRequests(&rowsFriendRequests[i]))
	}

	return out, nil
}
