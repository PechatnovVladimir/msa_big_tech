package social

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social/dto"
)

// Repository - интерфейс репозитория social
type Repository interface {
	Test() error
}

// UserService - интерфейс доступа к сервису пользователей
type UserService interface {
	Test() error
}

// Deps - зависимости
type Deps struct {
	SocialRepo  Repository
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
	RemoveFriend(ctx context.Context, in dto.RemoveFriendIN) (dto.RemoveFriendOUT, error)
	SendFriendRequest(ctx context.Context, in dto.SendFriendRequestIN) (dto.SendFriendRequestOUT, error)
}

var _ UseCase = (*Service)(nil)

func New(d Deps) *Service {
	return &Service{
		Deps: d,
	}
}
