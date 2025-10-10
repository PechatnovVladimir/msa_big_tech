package dto

import "time"

type SendMessageIN struct {
	ChatID string
	Text   string
}
type SendMessageOUT struct {
	MessageID string
	UserID    string
	ChatID    string
	Text      string
	CreatedAt time.Time
}
