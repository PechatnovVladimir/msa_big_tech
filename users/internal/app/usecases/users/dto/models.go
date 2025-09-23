package dto

type CreateProfileDTO struct {
	//Никнейм пользователя
	Nickname string
	//Биография пользователя
	Bio string
	//Ссылка на аватарку пользователя
	Avatar string
	//Пароль
	Password string
}

type UpdateProfileDTO struct {
	//ID пользователя
	ID string
	//Никнейм пользователя
	Nickname string
	//Биография пользователя
	Bio string
	//Ссылка на аватарку пользователя
	Avatar string
	//Пароль
	Password string
}

type SearchByNicknameDTO struct {
	query string
	limit int64
}
