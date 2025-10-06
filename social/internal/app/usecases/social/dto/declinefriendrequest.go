package dto

type DeclineFriendRequestIN struct {
	RequestID string
}
type DeclineFriendRequestOUT struct {
	RequestID string
	Status    StatusRequest
}
