package v1

import (
	"buf.build/go/protovalidate"
	"context"
	userPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) UpdateProfile(ctx context.Context, request *userPB.UpdateProfileRequest) (*userPB.UpdateProfileResponse, error) {
	//валидация по proto описанию
	err := protovalidate.Validate(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	data := dtoUpdateProfileFromUpdateProfileRequest(request)

	profile, err := s.uc.UpdateProfile(ctx, data)
	if err != nil {
		return nil, err
	}

	userProfile := responseUserProfileFromUserProfileModel(profile)

	return &userPB.UpdateProfileResponse{UserProfile: userProfile}, nil
}
