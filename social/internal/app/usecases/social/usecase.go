package social

import (
	"context"
	dtoRepo "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/repositories/social/dto"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
)

// Repository - интерфейс репозитория social
type Repository interface {
	SendFriendRequest(ctx context.Context, in dtoRepo.SendFriendRequestIN) error
	ListRequests(ctx context.Context, in dtoRepo.ListRequestsIN) (dtoRepo.ListRequestsOUT, error)
	GetFriendRequestByID(ctx context.Context, requestID string) (dtoRepo.FriendRequest, error)
	UpdateFriendRequest(ctx context.Context, in dtoRepo.FriendRequest) error
	ListFriends(ctx context.Context, in dtoRepo.ListFriendsIN) (dtoRepo.ListFriendsOUT, error)
	DeleteFriendRequest(ctx context.Context, in dtoRepo.DeleteFriendRequestIN) error
}

// UserService - интерфейс доступа к сервису пользователей
type UserService interface {
	Test() error
}

// AuthService - интерфейс доступа к сервису аутентификации
type AuthService interface {
	//GetAuthUser - получить аутентифицированного пользователя
	GetAuthUser() (string, error)
}

// Deps - зависимости
type Deps struct {
	SocialRepo  Repository
	AuthService AuthService
	UserService UserService
}

type Service struct {
	Deps
}

// UseCase - интерфейс сервиса чата
type UseCase interface {
	AcceptFriendRequest(ctx context.Context, in dto.AcceptFriendRequestIN) (dto.AcceptFriendRequestOUT, error)
	DeclineFriendRequest(ctx context.Context, in dto.DeclineFriendRequestIN) (dto.DeclineFriendRequestOUT, error)
	ListFriends(ctx context.Context, in dto.ListFriendsIN) (dto.ListFriendsOUT, error)
	ListRequests(ctx context.Context, in dto.ListRequestsIN) (dto.ListRequestsOUT, error)
	RemoveFriend(ctx context.Context, in dto.RemoveFriendIN) error
	SendFriendRequest(ctx context.Context, in dto.SendFriendRequestIN) (dto.SendFriendRequestOUT, error)
}

var _ UseCase = (*Service)(nil)

func New(d Deps) *Service {
	return &Service{
		Deps: d,
	}
}
