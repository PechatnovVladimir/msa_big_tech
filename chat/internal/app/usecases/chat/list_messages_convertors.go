package chat

import (
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"github.com/PechatnovVladimir/msa_big_tech/pkg/pagination"
)

func fromListMessagesIN(in *dto.ListMessagesIN) (chatID string, opts pagination.Options) {

	cursor := pagination.Cursor{
		ID:   &in.Cursor.ID,
		Time: &in.Cursor.Time,
	}

	opts = pagination.NewOptions(
		pagination.WithLimit(in.Limit),
		pagination.WithCursor(cursor),
	)

	return in.ChatID, opts
}

func toListMessageOUT(in []*chat.Message) *dto.ListMessagesOUT {

	cursor := dto.Cursor{}
	if len(in) != 0 {
		cursor.ID = in[len(in)-1].MessageID
		cursor.Time = in[len(in)-1].CreatedAt
	}

	out := &dto.ListMessagesOUT{
		Messages: in,
		Cursor:   &cursor,
	}

	return out
}
