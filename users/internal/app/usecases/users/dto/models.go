package dto

type CreateProfileDTO struct {
	//ID пользователя
	ID string
	//Никнейм пользователя
	Nickname string
	//Биография пользователя
	Bio string
	//Ссылка на аватарку пользователя
	Avatar string
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
}

type SearchByNicknameDTO struct {
	Query string
	Limit int64
}
