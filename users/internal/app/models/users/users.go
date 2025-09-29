package users

type UserProfile struct {
	//Идентификатор пользователя
	ID string
	//Никнейм пользователя
	Nickname string
	//Биография пользователя
	Bio string
	//Ссылка на аватарку пользователя
	Avatar string
}

func NewUserProfile() *UserProfile {
	return &UserProfile{}
}
