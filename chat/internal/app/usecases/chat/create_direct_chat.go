package chat

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
)

// CreateDirectChat - создание личного чата
func (s *Service) CreateDirectChat(ctx context.Context, in *dto.CreateDirectChatIN) (*dto.CreateDirectChatOUT, error) {

	data := fromCreateDirectChatIN(in)

	//запись в репозиторий
	chat, err := s.ChatRepo.CreateDirectChat(ctx, data)
	if err != nil {
		return &dto.CreateDirectChatOUT{}, err
	}

	out := toCreateDirectChatOUT(chat)
	//возвращаем ID созданного чата
	return out, nil
}
