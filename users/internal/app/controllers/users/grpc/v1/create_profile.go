package v1

import (
	"buf.build/go/protovalidate"
	"context"
	userPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) CreateProfile(ctx context.Context, request *userPB.CreateProfileRequest) (*userPB.CreateProfileResponse, error) {
	//валидация по proto описанию
	err := protovalidate.Validate(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	data := dtoCreateProfileFromCreateProfileRequest(request)

	profile, err := s.uc.CreateProfile(ctx, data)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	userProfile := responseUserProfileFromUserProfileModel(profile)

	return &userPB.CreateProfileResponse{UserProfile: userProfile}, nil
}
