package dto

type ListFriendsIN struct {
	UserID string
}
type ListFriendsOUT struct {
	FriendUserIDs []string
}
