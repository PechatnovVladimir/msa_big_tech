package social

import (
	"github.com/PechatnovVladimir/msa_big_tech/lib/pagination"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
	"time"
)

func fromListFriendsIN(in *dto.ListFriendsIN) (userID string, opts pagination.Options) {

	cursor := pagination.Cursor{
		ID:   &in.Cursor.ID,
		Time: &in.Cursor.Time,
	}

	if cursor.Time.Unix() == 0 {
		t := time.Now().UTC()
		cursor.Time = &t
	}

	opts = pagination.NewOptions(
		pagination.WithLimit(in.Limit),
		pagination.WithCursor(cursor),
	)

	return in.UserID, opts
}

func toListFriendsOUT(in []*social.Friend) *dto.ListFriendsOUT {
	if len(in) == 0 {
		return nil
	}

	userIDs := make([]string, len(in))
	for i, user := range in {
		userIDs[i] = user.ToUserID
	}
	return &dto.ListFriendsOUT{
		UserIDs: userIDs,
		Cursor: dto.Cursor{
			ID:   in[len(in)-1].FromUserID,
			Time: in[len(in)-1].CreatedAt,
		},
	}
}
