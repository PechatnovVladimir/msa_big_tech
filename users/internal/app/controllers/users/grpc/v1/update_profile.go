package v1

import (
	"context"
	userPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
)

func (s *Service) UpdateProfile(ctx context.Context, request *userPB.UpdateProfileRequest) (*userPB.UpdateProfileResponse, error) {
	data := dtoUpdateProfileFromUpdateProfileRequest(request)

	profile, err := s.uc.UpdateProfile(ctx, data)
	if err != nil {
		return nil, err
	}

	userProfile := responseUserProfileFromUserProfileModel(profile)

	return &userPB.UpdateProfileResponse{UserProfile: userProfile}, nil
}
