package chat

import (
	"context"
	dtoRepo "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/repositories/chat/dto"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"log"
)

func (s *Service) GetChat(ctx context.Context, in dto.GetChatIN) (dto.GetChatOUT, error) {
	log.Println("chat usecase called GetChat")

	//получаем текущего пользователя
	currentUser, ok := getCurrentUser(ctx)
	if !ok {
		return dto.GetChatOUT{}, dto.ErrUserNotAuthenticated
	}

	//получаем информацию о чате
	out, err := s.ChatRepo.GetChat(ctx, dtoRepo.GetChatIN{ChatID: in.ChatID})
	if err != nil {
		return dto.GetChatOUT{}, dto.ErrChatNotFound
	}

	//если текущий пользователь не владелец чата, то доступа нет
	if currentUser != out.UserID {
		return dto.GetChatOUT{}, dto.ErrChatPermissionDenied
	}

	return dto.GetChatOUT{
		ChatID: out.ChatID,
	}, nil
}
