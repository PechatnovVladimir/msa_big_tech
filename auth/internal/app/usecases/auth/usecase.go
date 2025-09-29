package auth

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/usecases/auth/dto"
)

type Repository interface {
	//SaveUser - сохранить пользователя
	SaveUser(ctx context.Context, userID string, email string, password string) error
	//CheckUserByEmail - проверить наличие пользователя
	CheckUserByEmail(ctx context.Context, email string) error
	//GetUserByEmailAndPassword - получить пользователя по email и проверить совпадение пароля
	GetUserByEmailAndPassword(ctx context.Context, email string, password string) (string, error)
	//SaveToken - сохранить токен
	SaveToken(ctx context.Context, userid string, accessToken string, refreshToken string) error
	//RefreshToken - перезаписать
	RefreshToken(ctx context.Context, userid string, accessToken string, refreshToken string) error
	//GetUserIDByRefreshToken - получить пользователя по refresh токену
	GetUserIDByRefreshToken(ctx context.Context, refreshToken string) (string, error)
}

// UserService - интерфейс доступа к сервису пользователей
type UserService interface {
	//CreateUserProfile - создать нового пользователя при регистрации
	CreateUserProfile(ctx context.Context, in dto.CreateUserProfileInDTO) (out dto.CreateUserProfileOutDTO, err error)
	//CheckUser - проверить наличие пользователя
	//CheckUser(ctx context.Context, userID string) (err error)
}

type Deps struct {
	AuthRepo    Repository
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

func cachePassword(password string) string {
	return fmt.Sprintf("%x", md5.Sum([]byte(password)))
}

func generateRandomToken(length int) (string, error) {
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b)[:length], nil
}
