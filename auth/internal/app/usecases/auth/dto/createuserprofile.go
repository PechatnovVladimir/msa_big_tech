package dto

type CreateUserProfileInDTO struct {
	UserID   string
	Nickname string
}

type CreateUserProfileOutDTO struct {
	UserID string
}
