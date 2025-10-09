package chat

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	dtoRepo "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/repositories/chat/dto"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"github.com/google/uuid"
)

// Repository - интерфейс репозитория чата
type Repository interface {
	//получить чат по user_id и participant_id между собеседниками
	CreateDirectChat(ctx context.Context, in *chat.Chat) (*chat.Chat, error)
	GetChatByUserAndParticipant(ctx context.Context, in dtoRepo.GetChatByUserAndParticipantIN) (out dtoRepo.GetChatByUserAndParticipantOUT, ok bool)
	GetChat(ctx context.Context, in dtoRepo.GetChatIN) (out dtoRepo.GetChatOUT, err error)
	ListChatMembers(ctx context.Context, in dtoRepo.ListChatMembersIN) (out dtoRepo.ListChatMembersOUT, err error)
	ListMessages(ctx context.Context, in dtoRepo.ListMessagesIN) (out dtoRepo.ListMessagesOUT, err error)
	ListUserChats(ctx context.Context, in dtoRepo.ListUserChatsIN) (out dtoRepo.ListUserChatsOUT, err error)
	SendMessage(ctx context.Context, in dtoRepo.SendMessageIN) (out dtoRepo.SendMessageOUT, err error)
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
	CreateDirectChat(ctx context.Context, in *dto.CreateDirectChatIN) (*dto.CreateDirectChatOUT, error)
	GetChat(ctx context.Context, in dto.GetChatIN) (dto.GetChatOUT, error)
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

func getCurrentUser(ctx context.Context) (string, bool) {
	//берем из контекста или ...
	//out, ok := ctx.Value("current_user").(string)
	//пока заглушка генерим на лету
	out := uuid.New().String()
	return out, true
}
