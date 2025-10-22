package v1

import (
	"context"
	userPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
)

func (s *Service) CreateProfile(ctx context.Context, request *userPB.CreateProfileRequest) (*userPB.CreateProfileResponse, error) {
	data := dtoCreateProfileFromCreateProfileRequest(request)

	profile, err := s.uc.CreateProfile(ctx, data)
	if err != nil {
		return nil, err
	}

	userProfile := responseUserProfileFromUserProfileModel(profile)

	return &userPB.CreateProfileResponse{UserProfile: userProfile}, nil
}
