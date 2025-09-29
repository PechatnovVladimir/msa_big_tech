package dto

type LoginInDTO struct {
	Email    string
	Password string
}
type LoginOutDTO struct {
	AccessToken  string
	RefreshToken string
	UserID       string
}
