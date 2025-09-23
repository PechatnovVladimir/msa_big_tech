package converter

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models"
	dtoRepo "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/repositories/social/dto"
	"time"
)

func SendFriendRequestDtoToRepo(ctx context.Context, in models.FriendRequest) (dtoRepo.SendFriendRequestIN, error) {
	return dtoRepo.SendFriendRequestIN{
		dtoRepo.FriendRequest{
			RequestID:  in.RequestID,
			FromUserID: in.FromUserID,
			ToUserID:   in.ToUserID,
			Status:     int(in.Status),
			CreateAt:   time.Now(),
		},
	}, nil
}
