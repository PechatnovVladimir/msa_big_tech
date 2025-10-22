package dto

type ListChatMembersIN struct {
	ChatID string
}
type ListChatMembersOUT struct {
	UserIDs []string
}
