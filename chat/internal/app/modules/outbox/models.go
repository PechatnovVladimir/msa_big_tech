package outbox

import (
	"time"
)

type AggregateType string

const (
	AggregateTypeMessageSent AggregateType = "message_sent"
)

type EventType string

const (
	EventTypeMessageSent EventType = "chat.message.sent"
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
