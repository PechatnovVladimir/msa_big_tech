package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Service) ListFriends(ctx context.Context, request *social.ListFriendsRequest) (*social.ListFriendsResponse, error) {

	data := fromListFriendsRequestToDto(request)

	log.Println("grps - ", data)

	listFriendsWithCursor, err := s.SocialUseCase.ListFriends(ctx, data)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	out := fromDtoToListFriendsResponse(listFriendsWithCursor)

	return out, nil
}
