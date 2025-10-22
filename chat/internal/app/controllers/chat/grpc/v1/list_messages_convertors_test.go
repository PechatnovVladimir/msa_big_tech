package v1

import (
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat/dto"
	"github.com/google/uuid"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	messages := make([]*chat.Message, 4)
	for i := 0; i < 4; i++ {
		messages[i] = &chat.Message{
			MessageID: uuid.New().String(),
			Text:      uuid.New().String(),
			SenderID:  uuid.New().String(),
			ChatID:    uuid.New().String(),
			CreatedAt: time.Now(),
		}
	}

	cursor := &dto.Cursor{
		ID:   uuid.New().String(),
		Time: time.Now(),
	}

	in := &dto.ListMessagesOUT{
		Messages: messages,
		Cursor:   cursor,
	}

	out := fromDtoToListMessagesResponse(in)

	for i, _ := range out.Messages {
		fmt.Println(out.Messages[i].MessageId)
	}

}
