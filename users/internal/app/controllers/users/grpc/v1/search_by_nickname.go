package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
	userPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
)

func (s *Service) SearchByNickname(ctx context.Context, request *userPB.SearchByNicknameRequest) (*userPB.SearchByNicknameResponse, error) {

	d := dto.SearchByNicknameDTO{
		Query: request.Query,
		Limit: request.Limit,
	}

	p, err := s.uc.SearchByNickname(ctx, d)
	if err != nil {
		return nil, err
	}

	userProfiles := make([]*userPB.UserProfile, len(p))

	for i, _ := range p {
		userProfiles[i] = &userPB.UserProfile{
			UserId:    p[i].ID,
			Nickname:  p[i].Nickname,
			AvatarUrl: p[i].Avatar,
			Bio:       p[i].Bio,
		}
	}

	return &userPB.SearchByNicknameResponse{UserProfile: userProfiles}, nil
}
