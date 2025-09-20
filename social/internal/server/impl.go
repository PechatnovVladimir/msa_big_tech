package server

import (
	"context"
	social "github.com/PechatnovVladimir/msa_big_tech/social/pkg/api"
	"log"
)

type SocialService struct {
	social.UnimplementedSocialServiceServer
}

func NewSocialService() *SocialService {
	return &SocialService{}
}

func (s *SocialService) SendFriendRequest(ctx context.Context, request *social.SendFriendRequestRequest) (*social.SendFriendRequestResponse, error) {
	log.Println("SocialService SendFriendRequest called")
	friendRequest := &social.FriendRequest{
		RequestId: 999,
		Status:    social.Status_PENDING,
	}
	return &social.SendFriendRequestResponse{FriendRequest: friendRequest}, nil
}

func (s *SocialService) ListRequests(ctx context.Context, request *social.ListRequestsRequest) (*social.ListRequestsResponse, error) {
	log.Println("SocialService ListRequests called")
	return nil, nil
}

func (s *SocialService) AcceptFriendRequest(ctx context.Context, request *social.AcceptFriendRequestRequest) (*social.AcceptFriendRequestResponse, error) {
	log.Println("SocialService AcceptFriendRequest called")
	return nil, nil
}

func (s *SocialService) DeclineFriendRequest(ctx context.Context, request *social.DeclineFriendRequestRequest) (*social.DeclineFriendRequestResponse, error) {
	log.Println("SocialService DeclineFriendRequest called")
	return nil, nil
}

func (s *SocialService) RemoveFriend(ctx context.Context, request *social.RemoveFriendRequest) (*social.RemoveFriendResponse, error) {
	log.Println("SocialService RemoveFriend called")
	return nil, nil
}

func (s *SocialService) ListFriends(ctx context.Context, request *social.ListFriendsRequest) (*social.ListFriendsResponse, error) {
	log.Println("SocialService ListFriends called")
	return nil, nil
}
