package social

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
)

func (s *Service) SendFriendRequest(ctx context.Context, in dto.SendFriendRequestIN) (dto.SendFriendRequestOUT, error) {
	return dto.SendFriendRequestOUT{}, nil
}
