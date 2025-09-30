package chat

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"log"
)

func (s *Service) CreateDirectChat(ctx context.Context, in dto.CreateDirectChatIN) (dto.CreateDirectChatOUT, error) {
	log.Println("usecase Chat Service CreateDirectChat called")

	//тестовый поход в репозиторий
	err := s.ChatRepo.Test()
	if err != nil {
		return dto.CreateDirectChatOUT{}, err
	}

	//тестовый поход в сервис UserService
	err = s.UserService.Test()
	if err != nil {
		return dto.CreateDirectChatOUT{}, err
	}

	return dto.CreateDirectChatOUT{}, nil
}
