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

func (s *Service) RemoveFriend(ctx context.Context, request *social.RemoveFriendRequest) (*social.RemoveFriendResponse, error) {
	log.Println("SocialService RemoveFriend called")
	//валидация по proto описанию
	err := protovalidate.Validate(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	_, _ = s.SocialUseCase.RemoveFriend(ctx, dto.RemoveFriendIN{})
	return nil, nil
}
