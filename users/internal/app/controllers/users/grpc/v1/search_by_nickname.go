package v1

import (
	"buf.build/go/protovalidate"
	"context"
	userPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) SearchByNickname(ctx context.Context, request *userPB.SearchByNicknameRequest) (*userPB.SearchByNicknameResponse, error) {
	//валидация по proto описанию

	err := protovalidate.Validate(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	data := dtoSearchByNicknameFromSearchByNicknameRequest(request)

	profiles, err := s.uc.SearchByNickname(ctx, data)
	if err != nil {
		return nil, err
	}

	userProfiles := responseSearchByNicknameFromUserProfilesModel(profiles)

	return &userPB.SearchByNicknameResponse{UserProfile: userProfiles}, nil
}
