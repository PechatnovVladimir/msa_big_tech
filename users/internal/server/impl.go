package server

import (
	"context"
	users "github.com/PechatnovVladimir/msa_big_tech/users/pkg/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserService struct {
	users.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) CreateProfile(ctx context.Context, request *users.CreateProfileRequest) (*users.CreateProfileResponse, error) {
	if request.UserId == 0 || request.Nickname == "" {
		return nil, status.Error(codes.InvalidArgument, "user id and nickname must be provided")
	}

	userProfile := users.UserProfile{
		UserId:   request.UserId,
		Nickname: request.Nickname,
	}

	return &users.CreateProfileResponse{UserProfile: &userProfile}, nil
}

func (s *UserService) UpdateProfile(ctx context.Context, request *users.UpdateProfileRequest) (*users.UpdateProfileResponse, error) {
	return nil, nil
}

func (s *UserService) GetProfileByID(ctx context.Context, request *users.GetProfileByIDRequest) (*users.GetProfileByIDResponse, error) {
	return nil, nil
}

func (s *UserService) GetProfileByNickname(ctx context.Context, request *users.GetProfileByNicknameRequest) (*users.GetProfileByNicknameResponse, error) {
	return nil, nil
}

func (s *UserService) SearchByNickname(ctx context.Context, request *users.SearchByNicknameRequest) (*users.SearchByNicknameResponse, error) {

	_ = request.Query

	userProfile := []*users.UserProfile{
		{UserId: 1,
			Nickname:  "pvv1",
			Bio:       "pvv1",
			AvatarUrl: "pvv1"},
		{UserId: 2,
			Nickname:  "pvv2",
			Bio:       "pvv2",
			AvatarUrl: "pvv2"},
	}

	return &users.SearchByNicknameResponse{UserProfile: userProfile}, nil
}
