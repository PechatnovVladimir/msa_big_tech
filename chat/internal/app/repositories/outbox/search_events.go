package outbox

import (
	"context"
	"github.com/Masterminds/squirrel"
	appoutbox "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/modules/outbox"
	"log"
	"time"
)

// SearchEvents выбирает события из outbox по заданным опциям.
// Возвращает пустой срез при ошибке (сигнатура без error).

func (r *Repository) SearchEvents(ctx context.Context, opts ...appoutbox.SearchEventsOption) []*appoutbox.Event {
	o := appoutbox.CollectSearchEventsOptions(opts...)

	// Базовый селект
	qb := r.sb.
		Select(tableOutboxEventsColumns...).
		From(tableOutboxEvents).
		OrderBy(columnOutboxCreatedAt).
		Limit(uint64(o.Limit))

	// Фильтры
	if o.OnlyUnpublished {
		qb = qb.Where(squirrel.Eq{columnOutboxPublishedAt: nil}) // IS NULL
	}
	// retry_count <= MaxRetryCount
	qb = qb.Where(squirrel.LtOrEq{columnOutboxRetryCount: o.MaxRetryCount})

	if o.AggregateType != nil {
		qb = qb.Where(squirrel.Eq{columnOutboxAggType: string(*o.AggregateType)})
	}
	if o.EventType != nil {
		qb = qb.Where(squirrel.Eq{columnOutboxEventType: string(*o.EventType)})
	}
	if o.NotBefore != nil {
		qb = qb.Where(squirrel.GtOrEq{columnOutboxCreatedAt: *o.NotBefore})
	}
	if o.NotAfter != nil {
		qb = qb.Where(squirrel.LtOrEq{columnOutboxCreatedAt: *o.NotAfter})
	}
	if o.DueAt != nil {
		// next_attempt_at IS NULL OR next_attempt_at <= dueAt
		qb = qb.Where(
			squirrel.Or{
				squirrel.Eq{columnOutboxNextAttemptAt: nil},
				squirrel.LtOrEq{columnOutboxNextAttemptAt: *o.DueAt},
			},
		)
	}

	// Блокировка строк для конкурентных воркеров
	if o.WithLock {
		qb = qb.Suffix("FOR UPDATE SKIP LOCKED")
	}

	// Выполнение
	conn := r.db.GetQueryEngine(ctx)

	//log.Println(qb.ToSql())

	var rows []OutboxEvent
	if err := conn.Selectx(ctx, &rows, qb); err != nil {
		log.Println(err.Error())
		return nil
	}

	// Маппинг в доменную модель
	events := make([]*appoutbox.Event, 0, len(rows))
	for i := range rows {
		var publishedAt *time.Time
		if rows[i].PublishedAt.Valid {
			t := rows[i].PublishedAt.V
			publishedAt = &t
		}
		var nextAttemptAt *time.Time
		if rows[i].NextAttemptAt.Valid {
			t := rows[i].NextAttemptAt.V
			nextAttemptAt = &t
		}
		events = append(events, &appoutbox.Event{
			ID:            rows[i].ID,
			AggregateType: appoutbox.AggregateType(rows[i].AggregateType),
			AggregateID:   rows[i].AggregateID,
			EventType:     appoutbox.EventType(rows[i].EventType),
			Payload:       rows[i].Payload,
			CreatedAt:     rows[i].CreatedAt,
			PublishedAt:   publishedAt,
			RetryCount:    rows[i].RetryCount,
			NextAttemptAt: nextAttemptAt,
		})
	}

	log.Println(events)

	return events
}
