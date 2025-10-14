package chat

import "time"

// Chat - чат
type Chat struct {
	//ChatID - ID чата в формате UUID
	ChatID string
	//Участники чата
	Members []string
	//Время создания
	CreateAt time.Time
}
