package social

import (
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
	"time"
)

type FriendsRow struct {
	UserID       string    `db:"user_id"`
	FriendUserID string    `db:"friend_user_id"`
	CreatedAt    time.Time `db:"created_at"`
}

func (r *FriendsRow) Values() []any {
	return []any{r.UserID, r.FriendUserID, r.CreatedAt}
}

func fromModelForCreateFriend(in *social.FriendRequest) *FriendsRow {
	if in == nil {
		return nil
	}
	return &FriendsRow{
		UserID:       in.FromUserID,
		FriendUserID: in.ToUserID,
		CreatedAt:    in.CreatedAt,
	}
}
