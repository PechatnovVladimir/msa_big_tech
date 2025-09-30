package social

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
	"log"
)

func (s *Service) ListFriends(ctx context.Context, in dto.ListFriendsIN) (dto.ListFriendsOUT, error) {
	log.Println("ListFriends")
	return dto.ListFriendsOUT{}, nil
}
