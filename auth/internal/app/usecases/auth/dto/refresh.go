package dto

type RefreshInDTO struct {
	RefreshToken string
}

type RefreshOutDTO struct {
	AccessToken  string
	RefreshToken string
	UserID       string
}
