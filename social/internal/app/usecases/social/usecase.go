package social

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
}

var _ UseCase = (*Service)(nil)

func New(d Deps) *Service {
	return &Service{
		Deps: d,
	}
}
