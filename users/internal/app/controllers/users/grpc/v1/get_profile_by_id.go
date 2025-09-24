package v1

import (
	"context"
	userPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
)

func (s *Service) GetProfileByID(ctx context.Context, request *userPB.GetProfileByIDRequest) (*userPB.GetProfileByIDResponse, error) {

	profileID := request.UserId

	p, err := s.uc.GetProfileByID(ctx, profileID)
	if err != nil {
		return nil, err
	}

	userProfile := &userPB.UserProfile{
		UserId:    p.ID.String(),
		Nickname:  p.Nickname,
		Bio:       p.Bio,
		AvatarUrl: p.Avatar,
	}

	return &userPB.GetProfileByIDResponse{UserProfile: userProfile}, nil

}
