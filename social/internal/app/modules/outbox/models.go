package outbox

import (
	"time"
)

type AggregateType string

const (
	AggregateTypeFriendRequest AggregateType = "friend_request"
)

type EventType string

const (
	EventTypeFriendRequest EventType = "social.friend.request"
	EventTypeFriendUpdated EventType = "social.friend.updated"
)

type Event struct {
	ID            string
	AggregateType AggregateType
	AggregateID   string
	EventType     EventType
	Payload       []byte
	CreatedAt     time.Time
	PublishedAt   *time.Time
	RetryCount    int
	NextAttemptAt *time.Time
}
