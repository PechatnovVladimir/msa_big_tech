package outbox

import (
	"github.com/google/uuid"
	"time"
)

const (
	AggregateTypeFriendRequest = "friend_request"
)

const (
	EventTypeFriendRequest = "social.friend.request"
	EventTypeFriendUpdated = "social.friend.updated"
)

type Event struct {
	ID            uuid.UUID
	AggregateType string
	AggregateID   string
	EventType     string
	Payload       []byte
	CreatedAt     time.Time
	PublishedAt   *time.Time
	RetryCount    int
}
