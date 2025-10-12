package v1

import (
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
)

func fromListChatMembersRequestToDto(request *chat.ListChatMembersRequest) *dto.ListChatMembersIN {
	return &dto.ListChatMembersIN{
		ChatID: request.ChatId,
	}
}

func fromDtoToListChatMembersResponse(in *dto.ListChatMembersOUT) *chat.ListChatMembersResponse {
	userIDs := make([]string, len(in.UserIDs))
	for i, userID := range in.UserIDs {
		userIDs[i] = *userID
	}

	return &chat.ListChatMembersResponse{
		UserIds: userIDs,
	}
}
