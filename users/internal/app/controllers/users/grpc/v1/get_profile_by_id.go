package v1

import (
	"context"
	userPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
)

func (s *Service) GetProfileByID(ctx context.Context, request *userPB.GetProfileByIDRequest) (*userPB.GetProfileByIDResponse, error) {
	data := dtoGetProfileByIDFromGetProfileByIDRequest(request)

	profile, err := s.uc.GetProfileByID(ctx, data)
	if err != nil {
		return nil, err
	}

	userProfile := responseUserProfileFromUserProfileModel(profile)

	return &userPB.GetProfileByIDResponse{UserProfile: userProfile}, nil

}
