package users

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
)

type UseCase interface {
	//CreateProfile создать профиль
	CreateProfile(ctx context.Context, dto dto.CreateProfile) (*users.UserProfile, error)
	//UpdateProfile обновить профиль
	UpdateProfile(ctx context.Context, dto dto.UpdateProfile) (*users.UserProfile, error)
	//GetProfileByID получить профиль по ID
	GetProfileByID(ctx context.Context, dto dto.GetProfileById) (*users.UserProfile, error)
	//GetProfileByNickname по иск по нику
	GetProfileByNickname(ctx context.Context, dto dto.GetProfileByNickname) (*users.UserProfile, error)
	//SearchByNickname поиск пользователей
	SearchByNickname(ctx context.Context, dto dto.SearchByNickname) ([]*users.UserProfile, error)
}

type Repository interface {
	CreateProfile(ctx context.Context, userProfile *users.UserProfile) (*users.UserProfile, error)
	UpdateProfile(ctx context.Context, userProfile *users.UserProfileForUpdate) (*users.UserProfile, error)
	DeleteProfile(ctx context.Context, profileID string) error
	GetProfileByID(ctx context.Context, profileID string) (*users.UserProfile, error)
	GetProfileByNickname(ctx context.Context, nickname string) (*users.UserProfile, error)
	SearchByNickname(ctx context.Context, filter *users.UserProfileFilter, limit *uint64) ([]*users.UserProfile, error)
}

type UserProvider interface {
	GetUserFromContext(ctx context.Context) (string, error)
}

type Deps struct {
	UserRepo     Repository
	UserProvider UserProvider
}

type Service struct {
	Deps
}

var _ UseCase = (*Service)(nil)

func New(d Deps) *Service {
	return &Service{
		Deps: d,
	}
}
