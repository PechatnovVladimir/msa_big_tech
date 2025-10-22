package auth

type Token struct {
	AccessToken  string
	RefreshToken string
	UserID       string
}
