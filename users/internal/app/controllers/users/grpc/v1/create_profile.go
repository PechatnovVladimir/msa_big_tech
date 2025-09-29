package v1

import (
	"buf.build/go/protovalidate"
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
	userPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) CreateProfile(ctx context.Context, request *userPB.CreateProfileRequest) (*userPB.CreateProfileResponse, error) {
	//валидация по proto описанию
	err := protovalidate.Validate(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	d := dto.CreateProfileDTO{
		Nickname: request.Nickname,
	}

	if request.Bio != nil {
		d.Bio = *request.Bio
	}

	if request.AvatarUrl != nil {
		d.Avatar = *request.AvatarUrl
	}

	p, err := s.uc.CreateProfile(ctx, d)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	userProfile := &userPB.UserProfile{
		UserId:    p.ID,
		Nickname:  p.Nickname,
		Bio:       p.Bio,
		AvatarUrl: p.Avatar,
	}

	return &userPB.CreateProfileResponse{UserProfile: userProfile}, nil
}
