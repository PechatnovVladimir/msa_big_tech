package dto

import "time"

type FriendRequest struct {
	RequestID  string
	FromUserID string
	ToUserID   string
	Status     int
	CreateAt   time.Time
}
