package chat

import (
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"github.com/google/uuid"
	"testing"
	"time"
)

func Test_toListMessageOUT(t *testing.T) {
	in := make([]*chat.Message, 2)
	for i := 0; i < 2; i++ {
		in[i] = &chat.Message{
			MessageID: uuid.New().String(),
			CreatedAt: time.Now(),
			SenderID:  uuid.New().String(),
			ChatID:    uuid.New().String(),
			Text:      uuid.New().String(),
		}
	}

	out := toListMessageOUT(in)

	fmt.Println(out.Messages[1].MessageID)
}
