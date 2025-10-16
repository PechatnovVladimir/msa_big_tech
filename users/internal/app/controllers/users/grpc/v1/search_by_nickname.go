package v1

import (
	"context"
	userPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
)

func (s *Service) SearchByNickname(ctx context.Context, request *userPB.SearchByNicknameRequest) (*userPB.SearchByNicknameResponse, error) {
	data := dtoSearchByNicknameFromSearchByNicknameRequest(request)

	profiles, err := s.uc.SearchByNickname(ctx, data)
	if err != nil {
		return nil, err
	}

	userProfiles := responseSearchByNicknameFromUserProfilesModel(profiles)

	return &userPB.SearchByNicknameResponse{UserProfile: userProfiles}, nil
}
