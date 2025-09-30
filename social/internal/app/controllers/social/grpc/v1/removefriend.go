package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
	"log"
)

func (s *Service) RemoveFriend(ctx context.Context, request *social.RemoveFriendRequest) (*social.RemoveFriendResponse, error) {
	log.Println("SocialService RemoveFriend called")
	return nil, nil
}
