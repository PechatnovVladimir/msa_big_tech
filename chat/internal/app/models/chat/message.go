package chat

import "time"

// Message - сообщения в чатах
type Message struct {
	//MessageID - ID сообщения
	MessageID string
	//ChatID - ID чата
	ChatID string
	//SenderID - ID собеседника
	SenderID string
	//Text - текст сообщения
	Text string
	//CreateAt - время создания сообщения
	CreatedAt time.Time
}

type Messages []Message
