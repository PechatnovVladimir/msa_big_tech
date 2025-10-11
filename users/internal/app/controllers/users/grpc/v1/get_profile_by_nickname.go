package v1

import (
	"context"
	userPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
)

func (s *Service) GetProfileByNickname(ctx context.Context, request *userPB.GetProfileByNicknameRequest) (*userPB.GetProfileByNicknameResponse, error) {
	nickname := request.Nickname

	p, err := s.uc.GetProfileByNickname(ctx, nickname)
	if err != nil {
		return nil, err
	}

	userProfile := &userPB.UserProfile{
		UserId:    p.ID,
		Nickname:  p.Nickname,
		Bio:       p.Bio,
		AvatarUrl: p.Avatar,
	}

	return &userPB.GetProfileByNicknameResponse{UserProfile: userProfile}, nil

}
