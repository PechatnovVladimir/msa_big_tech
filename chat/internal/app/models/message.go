package models

import "time"

// Message - сообщения в чатах
type Message struct {
	//MessageID - ID сообщения
	MessageID string
	//CreateAt - время создания сообщения
	CreatedAt time.Time
	//Text - текст сообщения
	Text string
	//ChatID - ID чата
	ChatID string
}

type Messages []Message
