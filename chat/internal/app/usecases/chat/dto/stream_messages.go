package dto

import (
	"time"
)

type StreamMessagesIN struct {
	ChatID           string
	SinceMessageTime time.Time
}
