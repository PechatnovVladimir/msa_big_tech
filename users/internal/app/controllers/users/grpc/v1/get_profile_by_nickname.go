package v1

import (
	"context"
	userPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
)

func (s *Service) GetProfileByNickname(ctx context.Context, request *userPB.GetProfileByNicknameRequest) (*userPB.GetProfileByNicknameResponse, error) {
	data := dtoGetProfileByNicknameFromGetProfileByNicknameRequest(request)

	profile, err := s.uc.GetProfileByNickname(ctx, data)
	if err != nil {
		return nil, err
	}

	userProfile := responseUserProfileFromUserProfileModel(profile)

	return &userPB.GetProfileByNicknameResponse{UserProfile: userProfile}, nil

}
