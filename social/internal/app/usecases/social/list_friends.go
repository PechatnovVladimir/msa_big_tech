package social

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
	"log"
)

func (s *Service) ListFriends(ctx context.Context, in *dto.ListFriendsIN) (*dto.ListFriendsOUT, error) {

	log.Println("UseCase - ", in)

	userID, paginationOpts := fromListFriendsIN(in)

	userIDs, err := s.SocialRepo.ListFriends(ctx, userID, paginationOpts)

	if err != nil {
		return nil, err
	}

	out := toListFriendsOUT(userIDs)

	return out, nil

}
