package social

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/repositories/social/dto"
	"log"
)

// SendFriendRequest - запись заявки в репозиторий
func (r *Repository) SendFriendRequest(ctx context.Context, in dto.SendFriendRequestIN) error {
	log.Println("Save friend request into DB...", in)
	return nil
}
