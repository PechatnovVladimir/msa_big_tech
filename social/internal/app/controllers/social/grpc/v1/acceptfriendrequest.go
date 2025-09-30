package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
	"log"
)

func (s *Service) AcceptFriendRequest(ctx context.Context, request *social.AcceptFriendRequestRequest) (*social.AcceptFriendRequestResponse, error) {
	log.Println("SocialService AcceptFriendRequest called")
	return nil, nil
}
