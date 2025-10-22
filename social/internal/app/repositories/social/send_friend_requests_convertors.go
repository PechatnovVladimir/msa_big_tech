package social

import (
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
	"time"
)

type FriendRequestsRow struct {
	ID         string    `db:"id"`
	FromUserID string    `db:"from_user_id"`
	ToUserID   string    `db:"to_user_id"`
	Status     int       `db:"status"`
	CreatedAt  time.Time `db:"created_at"`
}

func (r *FriendRequestsRow) Values() []any {
	return []any{r.ID, r.FromUserID, r.ToUserID, r.Status, r.CreatedAt}
}

func fromModelForSendFriendRequests(r *social.FriendRequest) *FriendRequestsRow {
	if r == nil {
		return nil
	}
	return &FriendRequestsRow{
		ID:         r.RequestID,
		FromUserID: r.FromUserID,
		ToUserID:   r.ToUserID,
		Status:     int(r.Status),
		CreatedAt:  r.CreatedAt,
	}
}

func toModelForSendFriendRequests(r *FriendRequestsRow) *social.FriendRequest {
	return &social.FriendRequest{
		RequestID:  r.ID,
		FromUserID: r.FromUserID,
		ToUserID:   r.ToUserID,
		Status:     social.StatusFriendRequest(r.Status),
		CreatedAt:  r.CreatedAt,
	}
}
