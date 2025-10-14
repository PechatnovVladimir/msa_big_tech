package chat

import "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"

func fromListChatMembersIN(in *dto.ListChatMembersIN) string {
	return in.ChatID
}

func toListChatMembersOUT(in []string) *dto.ListChatMembersOUT {
	return &dto.ListChatMembersOUT{
		UserIDs: in,
	}
}
