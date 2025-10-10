package v1

import (
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
)

func fromListUserChatsRequestToDto(in *chat.ListUserChatsRequest) *dto.ListUserChatsIN {
	return &dto.ListUserChatsIN{
		UserID: in.UserId,
	}
}

func fromDtoToListUserChatsResponse(in *dto.ListUserChatsOUT) *chat.ListUserChatsResponse {

	chats := make([]*chat.Chat, len(in.Chats))

	for i, _ := range in.Chats {
		chats[i] = &chat.Chat{
			ChatId: in.Chats[i].ChatID,
		}
	}

	return &chat.ListUserChatsResponse{
		Chats: chats,
	}
}
