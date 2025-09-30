package chat

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
)

// Repository - интерфейс репозитория чата
type Repository interface {
	Test() error
}

// UserService - интерфейс доступа к сервису пользователей
type UserService interface {
	Test() error
}

// Deps - зависимости
type Deps struct {
	ChatRepo    Repository
	UserService UserService
}

type Service struct {
	Deps
}

// UseCase - интерфейс сервиса чата
type UseCase interface {
	CreateDirectChat(ctx context.Context, in dto.CreateDirectChatIN) (dto.CreateDirectChatOUT, error)
	GetChat(ctx context.Context, in dto.GetChatIN) (dto.GetChatOut, error)
	ListChatMembers(ctx context.Context, in dto.ListChatMembersIN) (dto.ListChatMembersOUT, error)
	ListMessages(ctx context.Context, in dto.ListMessagesIN) (dto.ListMessagesOUT, error)
	ListUserChats(ctx context.Context, in dto.ListUserChatsIN) (dto.ListUserChatsOUT, error)
	SendMessage(ctx context.Context, in dto.SendMessageIN) (dto.SendMessageOut, error)
	StreamMessages(ctx context.Context, in dto.StreamMessagesIN) (dto.StreamMessagesOUT, error)
}

var _ UseCase = (*Service)(nil)

func New(d Deps) *Service {
	return &Service{
		Deps: d,
	}
}
