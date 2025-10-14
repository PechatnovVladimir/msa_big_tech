package chat

import "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"

func toModelGetChat(r []ChatsMembersRow) *chat.Chat {

	members := make([]string, len(r))
	for i, m := range r {
		members[i] = m.UserID
	}

	return &chat.Chat{
		ChatID:   r[0].ChatID,
		CreateAt: r[0].CreatedAt,
		Members:  members,
	}
}
