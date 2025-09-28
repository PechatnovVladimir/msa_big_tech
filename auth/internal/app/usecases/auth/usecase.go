package auth

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/usecases/auth/dto"
)

// UserService - интерфейс доступа к сервису пользователей
type UserService interface {
	//CreateUser - создать нового пользователя при регистрации
	CreateUser(ctx context.Context, nickname string, email string, password string) (userID string, err error)
	//CheckUser - проверить наличие пользователя и проверить совпадение пароля
	CheckUser(ctx context.Context, nickname string, password string) (userID string, err error)
}

type Deps struct {
	UserService UserService
}

type Service struct {
	Deps
}

type UseCase interface {
	Register(ctx context.Context, in dto.RegisterInDTO) (out dto.RegisterOutDTO, err error)
	Login(ctx context.Context, in dto.LoginInDTO) (out dto.LoginOutDTO, err error)
	Refresh(ctx context.Context, in dto.RefreshInDTO) (out dto.RefreshOutDTO, err error)
}

var _ UseCase = (*Service)(nil)

func New(d Deps) *Service {
	return &Service{
		Deps: d,
	}
}
