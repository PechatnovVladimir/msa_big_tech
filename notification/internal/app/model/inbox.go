package model

import (
	"time"
)

const (
	StatusReceived   = "received"
	StatusFailed     = "failed"
	StatusProcessing = "processing"
	StatusProcessed  = "processed"
)

type Inbox struct {
	ID          string
	Topic       string
	Partition   int32
	Offset      int64
	Payload     []byte
	Status      string
	Attempts    int
	LastError   string
	ReceivedAt  time.Time
	ProcessedAt *time.Time
}
