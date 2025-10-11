package social

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/repositories/social/dto"
	"time"
)

func (r *Repository) ListRequests(ctx context.Context, in dto.ListRequestsIN) (dto.ListRequestsOUT, error) {
	out := dto.ListRequestsOUT{FriendRequests: []dto.FriendRequest{
		{
			RequestID:  "9393DF6B-B37C-4B27-A965-00042AE1F911",
			FromUserID: "9393DF6B-B37C-4B27-A965-00042AE1F912",
			ToUserID:   "9393DF6B-B37C-4B27-A965-00042AE1F913",
			Status:     1,
			CreateAt:   time.Time{},
		},
		{
			RequestID:  "95B0F0A3-2B4E-4EA4-BC90-00116E9D3391",
			FromUserID: "95B0F0A3-2B4E-4EA4-BC90-00116E9D3392",
			ToUserID:   "95B0F0A3-2B4E-4EA4-BC90-00116E9D3393",
			Status:     2,
			CreateAt:   time.Time{},
		},
		{
			RequestID:  "95B0F0A3-2B4E-4EA4-BC90-00116E9D3391",
			FromUserID: "95B0F0A3-2B4E-4EA4-BC90-00116E9D3392",
			ToUserID:   "95B0F0A3-2B4E-4EA4-BC90-00116E9D3393",
			Status:     3,
			CreateAt:   time.Time{},
		},
	}}
	return out, nil
}
