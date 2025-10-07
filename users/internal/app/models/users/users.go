package users

import "time"

type UserProfile struct {
	//Идентификатор пользователя
	ID string
	//e-mail пользователя
	Email string
	//Никнейм пользователя
	Nickname string
	//Биография пользователя
	Bio string
	//Ссылка на аватарку пользователя
	Avatar string
	//Время создания
	CreateAt time.Time
}

func NewUserProfile() *UserProfile {
	return &UserProfile{}
}
