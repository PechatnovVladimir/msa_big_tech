package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
	"log"
)

func (s *Service) SendFriendRequest(ctx context.Context, request *social.SendFriendRequestRequest) (*social.SendFriendRequestResponse, error) {
	log.Println("SocialService SendFriendRequest called")
	friendRequest := &social.FriendRequest{
		RequestId: "147E46FE-1E80-4A65-8BFD-0015B5517E72",
		Status:    social.Status_PENDING,
	}
	return &social.SendFriendRequestResponse{FriendRequest: friendRequest}, nil
}
