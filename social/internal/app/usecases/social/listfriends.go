package social

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
	"github.com/google/uuid"
	"log"
)

func (s *Service) ListFriends(ctx context.Context, in dto.ListFriendsIN) (dto.ListFriendsOUT, error) {

	log.Println("ListFriends")
	userIDs := []string{
		uuid.New().String(),
		uuid.New().String(),
		uuid.New().String(),
	}

	return dto.ListFriendsOUT{
		UserIDs: userIDs,
		Cursor: dto.Cursor{
			UserID:    in.Cursor.UserID,
			CreatedAt: in.Cursor.CreatedAt,
		},
	}, nil
}
