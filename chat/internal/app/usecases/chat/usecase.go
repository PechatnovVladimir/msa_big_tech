package chat

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"github.com/PechatnovVladimir/msa_big_tech/pkg/pagination"
)

// Repository - интерфейс репозитория чата
type Repository interface {
	CreateDirectChat(ctx context.Context, chat *chat.Chat) (*chat.Chat, error)
	GetChat(ctx context.Context, chatID string) (*chat.Chat, error)
	ListUserChats(ctx context.Context, userID string) ([]*chat.Chat, error)
	SendMessage(ctx context.Context, m *chat.Message) (*chat.Message, error)
	ListChatMembers(ctx context.Context, chatID string) ([]*string, error)
	ListMessages(ctx context.Context, chatID string, opts pagination.Options) ([]*chat.Message, error)
}

// UserService - интерфейс доступа к сервису пользователей
type UserProvider interface {
	GetUserFromContext(ctx context.Context) (*chat.User, error)
}

// Deps - зависимости
type Deps struct {
	ChatRepo     Repository
	UserProvider UserProvider
}

type Service struct {
	Deps
}

// UseCase - интерфейс сервиса чата
type UseCase interface {
	CreateDirectChat(ctx context.Context, in *dto.CreateDirectChatIN) (*dto.CreateDirectChatOUT, error)
	GetChat(ctx context.Context, in *dto.GetChatIN) (*dto.GetChatOUT, error)
	ListUserChats(ctx context.Context, in *dto.ListUserChatsIN) (*dto.ListUserChatsOUT, error)
	SendMessage(ctx context.Context, in *dto.SendMessageIN) (*dto.SendMessageOUT, error)
	ListChatMembers(ctx context.Context, in *dto.ListChatMembersIN) (*dto.ListChatMembersOUT, error)
	ListMessages(ctx context.Context, in *dto.ListMessagesIN) (*dto.ListMessagesOUT, error)

	StreamMessages(ctx context.Context, in dto.StreamMessagesIN) (dto.StreamMessagesOUT, error)
}

var _ UseCase = (*Service)(nil)

func New(d Deps) *Service {
	return &Service{
		Deps: d,
	}
}
