package dto

import (
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	"time"
)

type Cursor struct {
	ID   string
	Time time.Time
}

type ListMessagesIN struct {
	ChatID string
	Limit  uint64
	Cursor Cursor
}

type ListMessagesOUT struct {
	Messages []*chat.Message
	Cursor   *Cursor
}
