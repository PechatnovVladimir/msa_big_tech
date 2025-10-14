package chat

import (
	"context"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"github.com/google/uuid"
	"slices"
	"time"
)

const maxTextSize = 4096

func (s *Service) SendMessage(ctx context.Context, in *dto.SendMessageIN) (*dto.SendMessageOUT, error) {
	const api = "ChatService.SendMessage"

	currentUser, err := s.UserProvider.GetUserFromContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("%s. %s: %w", "Пользователь не аутентифицирован", api, chat.ErrUnauthenticated)
	}

	if len(in.Text) > maxTextSize {
		return nil, fmt.Errorf("%s: %w: max messsage size is %d", api, chat.ErrInvalidArgument, maxTextSize)
	}

	chatInfo, err := s.ChatRepo.GetChat(ctx, in.ChatID)

	if err != nil {
		return &dto.SendMessageOUT{}, fmt.Errorf("%s: %w", api, err)
	}

	userIDs := chatInfo.Members

	//если текущий пользователь не является участником чата, то доступ к чату запрещен
	if !slices.Contains(userIDs, currentUser.UserID) {
		return nil, fmt.Errorf("%s: %w", api, chat.ErrPermissionDenied)
	}

	data := fromSendMessageIN(in)

	data.MessageID = uuid.New().String()
	data.SenderID = currentUser.UserID
	data.CreatedAt = time.Now()

	message, err := s.ChatRepo.SendMessage(ctx, data)

	if err != nil {
		return &dto.SendMessageOUT{}, err
	}

	out := toSendMessageOUT(message)

	return out, nil
}
