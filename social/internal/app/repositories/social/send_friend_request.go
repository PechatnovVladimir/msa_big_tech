package social

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
	"strings"
)

// SendFriendRequest - запись заявки в репозиторий
func (r *Repository) SendFriendRequest(ctx context.Context, fr *social.FriendRequest) (*social.FriendRequest, error) {
	friendRequestRow := fromModelForSendFriendRequests(fr)

	query := r.sb.
		Insert(friendRequestsTable).
		Columns(friendRequestsTableColumns...).
		Values(friendRequestRow.Values()...).
		Suffix("RETURNING " + strings.Join(friendRequestsTableColumns, ", "))

	pool := r.db.GetQueryEngine(ctx)

	var outRow FriendRequestsRow

	err := pool.Getx(ctx, &outRow, query)
	if err != nil {
		return nil, err
	}

	return toModelForSendFriendRequests(&outRow), nil
}
