package social

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
	"log"
)

func (s *Service) RemoveFriend(ctx context.Context, in dto.RemoveFriendIN) (dto.RemoveFriendOUT, error) {
	log.Println("RemoveFriend")
	return dto.RemoveFriendOUT{}, nil
}
