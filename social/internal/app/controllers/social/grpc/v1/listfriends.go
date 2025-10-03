package v1

import (
	"buf.build/go/protovalidate"
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/controllers/social/grpc/v1/converter"
	"github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Service) ListFriends(ctx context.Context, request *social.ListFriendsRequest) (*social.ListFriendsResponse, error) {
	log.Println("SocialService ListFriends called")
	//валидация по proto описанию
	err := protovalidate.Validate(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	inUC, err := converter.ListFriendsRequestProtoToDto(ctx, request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	outUC, err := s.SocialUseCase.ListFriends(ctx, inUC)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	out, err := converter.ListFriendsResponseDtoToProto(ctx, &outUC)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	return out, nil
}
