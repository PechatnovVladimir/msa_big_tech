package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
	userPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
)

func (s *Service) UpdateProfile(ctx context.Context, request *userPB.UpdateProfileRequest) (*userPB.UpdateProfileResponse, error) {

	d := dto.UpdateProfileDTO{
		ID: request.UserId,
	}

	if request.Nickname != nil {
		d.Nickname = *request.Nickname
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
		Bio:       p.Bio,
		AvatarUrl: p.Avatar,
	}

	return &userPB.UpdateProfileResponse{UserProfile: userProfile}, nil

}
