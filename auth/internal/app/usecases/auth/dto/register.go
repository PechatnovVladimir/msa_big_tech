package dto

type RegisterInDTO struct {
	Email    string
	Password string
}

type RegisterOutDTO struct {
	UserID string
}
