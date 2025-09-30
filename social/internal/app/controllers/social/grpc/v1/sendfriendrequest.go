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

func (s *Service) SendFriendRequest(ctx context.Context, request *social.SendFriendRequestRequest) (*social.SendFriendRequestResponse, error) {
	log.Println("SocialService SendFriendRequest called")
	//валидация по proto описанию
	err := protovalidate.Validate(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	_, _ = s.SocialUseCase.SendFriendRequest(ctx, dto.SendFriendRequestIN{})
	return nil, nil
}
