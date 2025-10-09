package dto

type UpdateProfile struct {
	//ID пользователя
	ID string
	//Никнейм пользователя
	Nickname *string
	//Email пользователя
	Email *string
	//Биография пользователя
	Bio *string
	//Ссылка на аватарку пользователя
	Avatar *string
}
