package outbox

import (
	"context"
	"database/sql"
	"fmt"
	appoutbox "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/modules/outbox"
	"time"
)

func (r *Repository) SaveEvent(ctx context.Context, e *appoutbox.Event) error {
	const api = "outbox.Repository.SaveEvent"

	row := OutboxEvent{
		ID:            e.ID,
		AggregateType: string(e.AggregateType),
		AggregateID:   e.AggregateID,
		EventType:     string(e.EventType),
		Payload:       notnullJSON(e.Payload),
		CreatedAt:     e.CreatedAt,
		PublishedAt: func(t *time.Time) sql.Null[time.Time] {
			if e.PublishedAt != nil {
				return sql.Null[time.Time]{V: *e.PublishedAt, Valid: true}
			}
			return sql.Null[time.Time]{}
		}(&e.CreatedAt),
		RetryCount: e.RetryCount,
		NextAttemptAt: func(t *time.Time) sql.Null[time.Time] {
			if e.NextAttemptAt != nil {
				return sql.Null[time.Time]{V: *e.NextAttemptAt, Valid: true}
			}
			return sql.Null[time.Time]{}
		}(&e.CreatedAt),
	}

	qb := r.sb.Insert(tableOutboxEvents).
		Columns(tableOutboxEventsColumns...).
		Values(row.Values(tableOutboxEventsColumns...)...)

	conn := r.db.GetQueryEngine(ctx)
	if _, err := conn.Execx(ctx, qb); err != nil {
		return fmt.Errorf("%s: %w", api, err)
	}

	return nil
}

func notnullJSON(data []byte) []byte {
	if data == nil {
		return []byte("[]")
	}
	return data
}
