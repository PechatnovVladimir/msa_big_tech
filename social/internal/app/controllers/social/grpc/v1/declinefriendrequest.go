package v1

import (
	"buf.build/go/protovalidate"
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
	"github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Service) DeclineFriendRequest(ctx context.Context, request *social.DeclineFriendRequestRequest) (*social.DeclineFriendRequestResponse, error) {
	log.Println("SocialService DeclineFriendRequest called")
	//валидация по proto описанию
	err := protovalidate.Validate(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	inUC := dto.DeclineFriendRequestIN{RequestID: request.RequestId}

	outUC, err := s.SocialUseCase.DeclineFriendRequest(ctx, inUC)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	friendRequest := &social.FriendRequest{
		RequestId: outUC.RequestID,
		Status:    social.Status(outUC.Status),
	}

	return &social.DeclineFriendRequestResponse{FriendRequest: friendRequest}, nil
}
