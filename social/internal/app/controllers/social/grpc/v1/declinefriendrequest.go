package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
	"log"
)

func (s *Service) DeclineFriendRequest(ctx context.Context, request *social.DeclineFriendRequestRequest) (*social.DeclineFriendRequestResponse, error) {
	log.Println("SocialService DeclineFriendRequest called")
	return nil, nil
}
