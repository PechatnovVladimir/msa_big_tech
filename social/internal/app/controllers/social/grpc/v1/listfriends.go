package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
	"log"
)

func (s *Service) ListFriends(ctx context.Context, request *social.ListFriendsRequest) (*social.ListFriendsResponse, error) {
	log.Println("SocialService ListFriends called")
	return nil, nil
}
