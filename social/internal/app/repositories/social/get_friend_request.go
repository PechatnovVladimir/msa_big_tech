package social

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
)

func (r *Repository) GetFriendRequest(ctx context.Context, requestID string) (*social.FriendRequest, error) {
	query := r.sb.Select("*").
		From(friendRequestsTable).
		Where(sq.Eq{friendRequestsTableColumnID: requestID})

	pool := r.db.GetQueryEngine(ctx)

	var rowFriendRequest FriendRequestsRow

	err := pool.Getx(ctx, &rowFriendRequest, query)

	if err != nil {
		return nil, err
	}

	out := toModelForSendFriendRequests(&rowFriendRequest)

	return out, nil

}
