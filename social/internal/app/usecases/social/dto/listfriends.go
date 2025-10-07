package dto

import (
	"time"
)

type Cursor struct {
	UserID    string
	CreatedAt time.Time
}

type ListFriendsIN struct {
	UserID string
	Limit  uint64
	Cursor Cursor
}
type ListFriendsOUT struct {
	UserIDs []string
	Cursor  Cursor
}
