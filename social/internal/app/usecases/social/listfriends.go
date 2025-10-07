package social

import (
	"context"
	dtoRepo "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/repositories/social/dto"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
	"log"
)

func (s *Service) ListFriends(ctx context.Context, in dto.ListFriendsIN) (dto.ListFriendsOUT, error) {

	log.Println("ListFriends")
	userIDs, err := s.SocialRepo.ListFriends(ctx, dtoRepo.ListFriendsIN{UserID: in.UserID})
	if err != nil {
		return dto.ListFriendsOUT{}, err
	}

	return dto.ListFriendsOUT{
		UserIDs: userIDs.FriendUserIDs,
		Cursor: dto.Cursor{
			UserID:    in.Cursor.UserID,
			CreatedAt: in.Cursor.CreatedAt,
		},
	}, nil
}
