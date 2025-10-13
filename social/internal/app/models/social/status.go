package social

type StatusFriendRequest int

const (
	UNKNOWN StatusFriendRequest = iota
	PENDING
	ACCEPTED
	DECLINED
)
