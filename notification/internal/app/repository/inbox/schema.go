package inbox

import (
	"database/sql"
	"time"
)

const tableInboxMessages = "inbox_messages"

const (
	columnInboxID          = "id"
	columnInboxTopic       = "topic"
	columnInboxPartition   = "partition"
	columnInboxOffset      = "offsett"
	columnInboxPayload     = "payload"
	columnInboxStatus      = "status"
	columnInboxAttempts    = "attempts"
	columnInboxLastError   = "last_error"
	columnInboxReceivedAt  = "received_at"
	columnInboxProcessedAt = "processed_at"
)

var (
	tableInboxMessagesColumns = []string{
		columnInboxID,
		columnInboxTopic,
		columnInboxPartition,
		columnInboxOffset,
		columnInboxPayload,
		columnInboxStatus,
		columnInboxAttempts,
		columnInboxLastError,
		columnInboxReceivedAt,
		columnInboxProcessedAt,
	}
)

type Inbox struct {
	ID          string              `db:"id"`
	Topic       string              `db:"topic"`
	Partition   int32               `db:"partition"`
	Offset      int64               `db:"offsett"`
	Payload     []byte              `db:"payload"`
	Status      string              `db:"status"`
	Attempts    int                 `db:"attempts"`
	LastError   string              `db:"last_error"`
	ReceivedAt  time.Time           `db:"received_at"`
	ProcessedAt sql.Null[time.Time] `db:"processed_at"`
}

func (e *Inbox) mapFields() map[string]any {
	return map[string]any{
		columnInboxID:          e.ID,
		columnInboxTopic:       e.Topic,
		columnInboxPartition:   e.Partition,
		columnInboxOffset:      e.Offset,
		columnInboxPayload:     e.Payload,
		columnInboxStatus:      e.Status,
		columnInboxAttempts:    e.Attempts,
		columnInboxLastError:   e.LastError,
		columnInboxReceivedAt:  e.ReceivedAt,
		columnInboxProcessedAt: e.ProcessedAt,
	}
}

func (e *Inbox) Values(columns ...string) []any {
	m := e.mapFields()
	values := make([]any, 0, len(columns))
	for i := range columns {
		if v, ok := m[columns[i]]; ok {
			values = append(values, v)
		} else {
			values = append(values, nil)
		}
	}
	return values
}
