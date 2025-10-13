package social

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/models/social"
	dtoRepo "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/repositories/social/dto"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
)

// Repository - интерфейс репозитория social
type Repository interface {
	SendFriendRequest(ctx context.Context, fr *social.FriendRequest) (*social.FriendRequest, error)
	ListRequests(ctx context.Context, userID string) ([]*social.FriendRequest, error)

	GetFriendRequestByID(ctx context.Context, requestID string) (dtoRepo.FriendRequest, error)
	UpdateFriendRequest(ctx context.Context, in dtoRepo.FriendRequest) error
	ListFriends(ctx context.Context, in dtoRepo.ListFriendsIN) (dtoRepo.ListFriendsOUT, error)
	DeleteFriendRequest(ctx context.Context, in dtoRepo.DeleteFriendRequestIN) error
}

// UserProvider - интерфейс доступа к сервису пользователей
type UserProvider interface {
	GetUserFromContext(ctx context.Context) (*social.User, error)
}

// AuthProvider - интерфейс доступа к сервису аутентификации
type AuthProvider interface {
	//GetAuthUser - получить аутентифицированного пользователя
	GetAuthUser() (string, error)
}

// Deps - зависимости
type Deps struct {
	SocialRepo   Repository
	AuthProvider AuthProvider
	UserProvider UserProvider
}

type Service struct {
	Deps
}

// UseCase - интерфейс сервиса чата
type UseCase interface {
	SendFriendRequest(ctx context.Context, in *dto.SendFriendRequestIN) (*dto.SendFriendRequestOUT, error)
	ListRequests(ctx context.Context, in *dto.ListRequestsIN) (*dto.ListRequestsOUT, error)

	AcceptFriendRequest(ctx context.Context, in dto.AcceptFriendRequestIN) (dto.AcceptFriendRequestOUT, error)
	DeclineFriendRequest(ctx context.Context, in dto.DeclineFriendRequestIN) (dto.DeclineFriendRequestOUT, error)
	ListFriends(ctx context.Context, in dto.ListFriendsIN) (dto.ListFriendsOUT, error)
	RemoveFriend(ctx context.Context, in dto.RemoveFriendIN) error
}

var _ UseCase = (*Service)(nil)

func New(d Deps) *Service {
	return &Service{
		Deps: d,
	}
}
