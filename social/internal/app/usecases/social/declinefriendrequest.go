package social

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
	"log"
)

func (s *Service) DeclineFriendRequest(ctx context.Context, in dto.DeclineFriendRequestIN) (dto.DeclineFriendRequestOUT, error) {
	log.Println("DeclineFriendRequest")
	return dto.DeclineFriendRequestOUT{}, nil
}
