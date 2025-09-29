package dto

type CreateUserProfileInDTO struct {
	userID   string
	nickname string
}

type CreateUserProfileOutDTO struct {
	userID string
}
