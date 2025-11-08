package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
	"github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) RemoveFriend(ctx context.Context, request *social.RemoveFriendRequest) (*social.RemoveFriendResponse, error) {

	err := s.SocialUseCase.RemoveFriend(ctx, dto.RemoveFriendIN{
		UserID: request.UserId,
	})

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &social.RemoveFriendResponse{}, nil
}
