package dto

import "time"

type CreateProfileDTO struct {
	//ID пользователя
	ID string
	//Никнейм пользователя
	Nickname string
	//Email пользователя
	Email string
	//Биография пользователя
	Bio *string
	//Ссылка на аватарку пользователя
	Avatar *string
}

type UpdateProfileDTO struct {
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

type Query struct {
	IDs         []string
	Email       *string
	Nickname    *string
	CreatedFrom *time.Time
	CreatedTo   *time.Time
}

type SearchByNicknameDTO struct {
	Query Query
	Limit uint64
}
