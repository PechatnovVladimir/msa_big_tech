package dto

import "time"

type Cursor struct {
	ID   string
	Time time.Time
}

type ListFriendsIN struct {
	UserID string
	Limit  uint64
	Cursor Cursor
}
type ListFriendsOUT struct {
	FriendUserIDs []string
}
