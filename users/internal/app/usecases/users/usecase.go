package users

import (
	"context"
	"crypto/md5"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
)

type (
	UserRepository interface {
		CreateProfile(ctx context.Context, userProfile *users.UserProfile) error
		UpdateProfile(ctx context.Context, userProfile *users.UserProfile) error
		DeleteProfile(ctx context.Context, profileID string) error
		GetProfileByID(ctx context.Context, profileID string) (*users.UserProfile, error)
		GetProfileByNickname(ctx context.Context, nickname string) (*users.UserProfile, error)
		SearchByNickname(ctx context.Context, nickname string) ([]*users.UserProfile, error)
	}
)
type UserService struct {
	userRepo UserRepository
}

type UserUsecase interface {
	//CreateProfile создать профиль
	CreateProfile(ctx context.Context, dto dto.CreateProfileDTO) (*users.UserProfile, error)
	//UpdateProfile обновить профиль
	UpdateProfile(ctx context.Context, dto dto.UpdateProfileDTO) (*users.UserProfile, error)
	//GetProfileByID получить профиль по ID
	GetProfileByID(ctx context.Context, id string) (*users.UserProfile, error)
	//GetProfileByNickname по иск по нику
	GetProfileByNickname(ctx context.Context, nickname string) (*users.UserProfile, error)
	//SearchByNickname поиск пользователей
	SearchByNickname(ctx context.Context, dto dto.SearchByNicknameDTO) ([]*users.UserProfile, error)
}

var _ UserUsecase = (*UserService)(nil)

func NewUserService(userRepo UserRepository) *UserService {
	return &UserService{
		userRepo: userRepo,
	}
}

func cachePassword(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}
