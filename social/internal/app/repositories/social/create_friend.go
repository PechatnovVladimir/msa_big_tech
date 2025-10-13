package social

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"

)

func (r *Repository) CreateFriend(ctx context.Context, fr *social.FriendRequest) error {
	query:=r.sb.
		Insert(friendsTable).
		Columns(friendsTableColumns...).

	return nil
}
