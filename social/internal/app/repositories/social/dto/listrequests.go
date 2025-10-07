package dto

type ListRequestsIN struct {
	UserID string
}

type ListRequestsOUT struct {
	FriendRequests []FriendRequest
}
