package chat

import (
	"context"
	dtoRepo "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/repositories/chat/dto"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"github.com/google/uuid"
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
	_, ok = s.ChatRepo.GetChatByUserAndParticipant(ctx, dtoRepo.GetChatByUserAndParticipantIN{
		UserId:        currentUser,
		ParticipantId: in.ParticipantID,
	})
	if ok {
		return dto.CreateDirectChatOUT{}, dto.ErrChatAlreadyExists
	}

	chatDTO := dtoRepo.CreateDirectChatIN{
		ChatID:        uuid.New().String(),
		UserID:        currentUser,
		ParticipantID: in.ParticipantID,
	}

	//запись в репозиторий
	err := s.ChatRepo.CreateDirectChat(ctx, chatDTO)
	if err != nil {
		return dto.CreateDirectChatOUT{}, err
	}

	//возвращаем ID созданного чата
	return dto.CreateDirectChatOUT{
		ChatID: chatDTO.ChatID,
	}, nil
}
