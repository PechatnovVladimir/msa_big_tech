package v1

import (
	"buf.build/go/protovalidate"
	"context"
	userPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Service) GetProfileByNickname(ctx context.Context, request *userPB.GetProfileByNicknameRequest) (*userPB.GetProfileByNicknameResponse, error) {
	//валидация по proto описанию
	err := protovalidate.Validate(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	nickname := request.Nickname

	p, err := s.uc.GetProfileByNickname(ctx, nickname)
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

	return &userPB.GetProfileByNicknameResponse{UserProfile: userProfile}, nil

}
