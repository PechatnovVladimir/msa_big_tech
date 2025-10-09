package chat

// Chat - чат
type Chat struct {
	//ChatID - ID чата в формате UUID
	ChatID string
	//UserID - ID пользователя создавшего чат в формате UUID
	UserID string
}

type Chats []Chat
