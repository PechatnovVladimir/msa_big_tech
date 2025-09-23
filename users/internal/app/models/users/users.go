package users

import "github.com/google/uuid"

type UserID string

func (id UserID) String() string {
	return string(id)
}

type UserProfile struct {
	//Идентификатор пользователя
	ID UserID
	//Никнейм пользователя
	Nickname string
	//Биография пользователя
	Bio string
	//Ссылка на аватарку пользователя
	Avatar string
	//Пароль
	Password string
}

func NewUserProfile() *UserProfile {
	return &UserProfile{
		ID: UserID(uuid.New().String()),
	}
}
