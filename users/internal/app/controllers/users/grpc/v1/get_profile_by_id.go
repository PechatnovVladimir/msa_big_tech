package v1

import (
	"buf.build/go/protovalidate"
	"context"
	userPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetProfileByID(ctx context.Context, request *userPB.GetProfileByIDRequest) (*userPB.GetProfileByIDResponse, error) {
	//валидация по proto описанию
	err := protovalidate.Validate(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	profileID := request.UserId

	p, err := s.uc.GetProfileByID(ctx, profileID)
	if err != nil {
		return nil, err
	}

	userProfile := &userPB.UserProfile{
		UserId:    p.ID,
		Email:     p.Email,
		Nickname:  p.Nickname,
		Bio:       p.Bio,
		AvatarUrl: p.Avatar,
	}

	return &userPB.GetProfileByIDResponse{UserProfile: userProfile}, nil

}
