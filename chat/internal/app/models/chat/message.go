package chat

import "time"

// Message - сообщения в чатах
type Message struct {
	//MessageID - ID сообщения
	MessageID string `json:"message_id"`
	//ChatID - ID чата
	ChatID string `json:"chat_id"`
	//SenderID - ID собеседника
	SenderID string `json:"sender_id"`
	//Text - текст сообщения
	Text string `json:"text"`
	//CreateAt - время создания сообщения
	CreatedAt time.Time `json:"created_at"`
}

type Messages []Message
