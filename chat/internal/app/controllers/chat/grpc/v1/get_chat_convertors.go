package v1

import (
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
)

func fromGetChatRequestToDto(in *chat.GetChatRequest) *dto.GetChatIN {
	return &dto.GetChatIN{
		ChatID: in.ChatId,
	}
}

func fromDtoToGetChatResponse(in *dto.GetChatOUT) *chat.GetChatResponse {
	return &chat.GetChatResponse{
		Chat: &chat.Chat{
			ChatId: in.ChatID,
		},
	}
}
