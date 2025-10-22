package v1

import (
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	chatPB "github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
)

func modelMessageToChatPB(m *chat.Message) *chatPB.Message {
	if m == nil {
		return nil
	}

	return &chatPB.Message{
		MessageId: m.MessageID,
		Text:      m.Text,
	}
}
