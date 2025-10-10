package chat

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"log"
)

func (s *Service) GetChat(ctx context.Context, in *dto.GetChatIN) (*dto.GetChatOUT, error) {
	log.Println("chat usecase called GetChat")

	//получаем текущего пользователя
	//currentUser, ok := getCurrentUser(ctx)
	//if !ok {
	//	return dto.GetChatOUT{}, dto.ErrUserNotAuthenticated
	//}

	data := fromGetChatIN(in)

	//получаем информацию о чате
	chat, err := s.ChatRepo.GetChat(ctx, data)

	if err != nil {
		return &dto.GetChatOUT{}, dto.ErrChatNotFound
	}

	//если текущий пользователь не владелец чата, то доступа нет
	//if currentUser != chat.UserID {
	//	return dto.GetChatOUT{}, dto.ErrChatPermissionDenied
	//}
	return toGetChatOUT(chat), nil
}
