package models

type StatusFriendRequest int

const (
	UNKNOWN StatusFriendRequest = iota
	PENDING
	ACCEPTED
	DECLINED
)
