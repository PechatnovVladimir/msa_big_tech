package v1

import (
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
)

// конвертация request в dto
func fromCreateDirectChatRequestToDTO(request *chat.CreateDirectChatRequest) *dto.CreateDirectChatIN {
	out := &dto.CreateDirectChatIN{}
	out.ParticipantID = request.ParticipantId
	return out
}

// конвертация dto в response
func fromDtoToCreateDirectChatResponse(in *dto.CreateDirectChatOUT) *chat.CreateDirectChatResponse {
	out := &chat.CreateDirectChatResponse{
		ChatId: in.ChatID,
	}
	return out
}
