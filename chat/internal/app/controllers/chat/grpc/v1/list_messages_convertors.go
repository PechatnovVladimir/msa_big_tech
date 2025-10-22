package v1

import (
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func fromListMessagesRequest(in *chat.ListMessagesRequest) *dto.ListMessagesIN {

	cursor := &dto.Cursor{
		ID:   in.Cursor.MessageId,
		Time: in.Cursor.CreatedAt.AsTime(),
	}

	out := &dto.ListMessagesIN{
		ChatID: in.ChatId,
		Limit:  in.Limit,
		Cursor: *cursor,
	}

	return out
}

func fromDtoToListMessagesResponse(in *dto.ListMessagesOUT) *chat.ListMessagesResponse {

	nextCursor := &chat.Cursor{
		MessageId: in.Cursor.ID,
		CreatedAt: timestamppb.New(in.Cursor.Time),
	}

	messages := make([]*chat.Message, len(in.Messages))
	for i := range in.Messages {
		messageID := in.Messages[i].MessageID
		text := in.Messages[i].Text
		messages[i] = &chat.Message{
			MessageId: messageID,
			Text:      text,
		}
	}

	return &chat.ListMessagesResponse{
		Messages:   messages,
		NextCursor: nextCursor,
	}

}
