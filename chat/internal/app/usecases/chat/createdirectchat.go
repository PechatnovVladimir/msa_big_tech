package chat

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"log"
)

// CreateDirectChat - создание личного чата
func (s *Service) CreateDirectChat(ctx context.Context, in dto.CreateDirectChatIN) (dto.CreateDirectChatOUT, error) {
	log.Println("usecase Chat Service CreateDirectChat called")
	//получаем текущего пользователя
	currentUser, ok := getCurrentUser(ctx)
	if !ok {
		return dto.CreateDirectChatOUT{}, dto.ErrUserNotAuthenticated
	}

	//проверяем наличие чата между currentUser и participant_id - поход в репозиторий
	_, ok = s.ChatRepo.GetChatByUserAndParticipant(ctx, currentUser, in.ParticipantID)
	if ok {
		return dto.CreateDirectChatOUT{}, dto.ErrChatAlreadyExists
	}

	//chat := models.Chat{
	//	ChatID:        uuid.New().String(),
	//	UserID:        currentUser,
	//	ParticipantID: in.ParticipantID,
	//}

	return dto.CreateDirectChatOUT{}, nil
}

func getCurrentUser(ctx context.Context) (string, bool) {
	//берем из контекста или ...
	out, ok := ctx.Value("current_user").(string)
	return out, ok
}
