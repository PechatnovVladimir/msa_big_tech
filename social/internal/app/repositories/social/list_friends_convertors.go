package social

import (
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
	"time"
)

type FriendRow struct {
	UserID       string    `db:"user_id"`
	FriendUserID string    `db:"friend_user_id"`
	CreatedAt    time.Time `db:"created_at"`
}

func toModelForListFriends(m FriendRow) *social.Friend {
	return &social.Friend{
		FromUserID: m.UserID,
		ToUserID:   m.FriendUserID,
		CreatedAt:  m.CreatedAt,
	}
}
