package social

import (
	"context"
	dtoRepo "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/repositories/social/dto"
	"github.com/google/uuid"
)

func (r *Repository) ListFriends(ctx context.Context, in dtoRepo.ListFriendsIN) (dtoRepo.ListFriendsOUT, error) {
	return dtoRepo.ListFriendsOUT{
		[]string{uuid.New().String(), uuid.New().String(), uuid.New().String()},
	}, nil
}
