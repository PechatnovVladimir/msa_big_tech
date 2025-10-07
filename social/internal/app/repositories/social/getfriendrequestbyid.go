package social

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/repositories/social/dto"
	"github.com/google/uuid"
	"time"
)

func (r *Repository) GetFriendRequestByID(ctx context.Context, requestID string) (dto.FriendRequest, error) {
	return dto.FriendRequest{
		RequestID:  requestID,
		FromUserID: uuid.New().String(),
		ToUserID:   uuid.New().String(),
		CreateAt:   time.Now(),
		Status:     1,
	}, nil
}
