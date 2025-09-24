package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
	userPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) CreateProfile(ctx context.Context, request *userPB.CreateProfileRequest) (*userPB.CreateProfileResponse, error) {
	if request.Password == "" || request.Nickname == "" {
		return nil, status.Error(codes.InvalidArgument, "nickname and password must be provided")
	}

	d := dto.CreateProfileDTO{
		Nickname: request.Nickname,
		Bio:      *request.Bio,
		Avatar:   *request.AvatarUrl,
		Password: request.Password,
	}

	p, err := s.uc.CreateProfile(ctx, d)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	userProfile := &userPB.UserProfile{
		UserId:    p.ID.String(),
		Nickname:  p.Nickname,
		Bio:       p.Bio,
		AvatarUrl: p.Avatar,
	}

	return &userPB.CreateProfileResponse{UserProfile: userProfile}, nil
}
