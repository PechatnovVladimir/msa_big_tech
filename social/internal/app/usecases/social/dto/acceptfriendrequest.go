package dto

type AcceptFriendRequestIN struct {
	RequestID string
}
type AcceptFriendRequestOUT struct {
	RequestID string
	Status    StatusRequest
}
