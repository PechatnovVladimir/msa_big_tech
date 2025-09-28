package v1

import (
	"buf.build/go/protovalidate"
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
	userPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) UpdateProfile(ctx context.Context, request *userPB.UpdateProfileRequest) (*userPB.UpdateProfileResponse, error) {

	err := protovalidate.Validate(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	d := dto.UpdateProfileDTO{
		ID: request.UserId,
	}

	if request.Nickname != nil {
		d.Nickname = *request.Nickname
	}

	if request.Email != nil {
		d.Email = *request.Email
	}

	if request.Bio != nil {
		d.Bio = *request.Bio
	}

	if request.AvatarUrl != nil {
		d.Avatar = *request.AvatarUrl
	}

	if request.Password != nil {
		d.Password = *request.Password
	}

	p, err := s.uc.UpdateProfile(ctx, d)
	if err != nil {
		return nil, err
	}

	userProfile := &userPB.UserProfile{
		UserId:    p.ID.String(),
		Nickname:  p.Nickname,
		Email:     p.Email,
		Bio:       p.Bio,
		AvatarUrl: p.Avatar,
	}

	return &userPB.UpdateProfileResponse{UserProfile: userProfile}, nil

}
