package dto

type GetChatIN struct {
	ChatID string
}
type GetChatOUT struct {
	ChatID        string
	UserID        string
	ParticipantID string
}
