package models

// Chat - чат
type Chat struct {
	//ChatID - ID чата в формате UUID
	ChatID string
	//UserID - ID пользователя создавшего чат в формате UUID
	UserID string
	//ParticipantID - ID пользователя собеседника в формате UUID
	ParticipantID string
}

type Chats []Chat
