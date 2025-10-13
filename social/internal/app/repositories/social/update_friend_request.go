package social

import (
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
	"strings"
)

func (r *Repository) ChangeStatusFriendRequest(ctx context.Context, requestID string, status int) (*social.FriendRequest, error) {
	query := r.sb.Update(friendRequestsTable).
		Set("status", status).
		Where(sq.Eq{"request_id": requestID}).
		Suffix("RETURNING " + strings.Join(friendRequestsTableColumns, ","))

	pool := r.db.GetQueryEngine(ctx)

	var outRow FriendRequestsRow
	if err := pool.Getx(ctx, &outRow, query); err != nil {
		return &social.FriendRequest{}, err
	}

	return &social.FriendRequest{
		RequestID:  outRow.ID,
		FromUserID: outRow.FromUserID,
		ToUserID:   outRow.ToUserID,
		Status:     social.StatusFriendRequest(outRow.Status),
		CreatedAt:  outRow.CreatedAt,
	}, nil
}
