package server

import (
	"context"
	social "github.com/PechatnovVladimir/msa_big_tech/social/pkg/api"
)

type SocialService struct {
	social.UnimplementedSocialServiceServer
}

func NewSocialService() *SocialService {
	return &SocialService{}
}

func (s *SocialService) SendFriendRequest(ctx context.Context, request *social.SendFriendRequestRequest) (*social.SendFriendRequestResponse, error) {
	friendRequest := &social.FriendRequest{
		RequestId: 999,
		Status:    social.Status_PENDING,
	}
	return &social.SendFriendRequestResponse{FriendRequest: friendRequest}, nil
}

func (s *SocialService) ListRequests(ctx context.Context, request *social.ListRequestsRequest) (*social.ListRequestsResponse, error) {
	return nil, nil
}

func (s *SocialService) AcceptFriendRequest(ctx context.Context, request *social.AcceptFriendRequestRequest) (*social.AcceptFriendRequestResponse, error) {
	return nil, nil
}

func (s *SocialService) DeclineFriendRequest(ctx context.Context, request *social.DeclineFriendRequestRequest) (*social.DeclineFriendRequestResponse, error) {
	return nil, nil
}

func (s *SocialService) RemoveFriend(ctx context.Context, request *social.RemoveFriendRequest) (*social.RemoveFriendResponse, error) {
	return nil, nil
}

func (s *SocialService) ListFriends(ctx context.Context, request *social.ListFriendsRequest) (*social.ListFriendsResponse, error) {
	return nil, nil
}
