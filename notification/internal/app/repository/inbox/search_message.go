package inbox

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/PechatnovVladimir/msa_big_tech/notification/internal/app/model"
	"github.com/PechatnovVladimir/msa_big_tech/notification/internal/app/usecase/inbox"
	"log"
	"time"
)

func (r *Repository) SearchMessage(ctx context.Context, opts ...inbox.SearchMessageOption) []*model.Inbox {
	o := inbox.CollectSearchMessageOptions(opts...)

	// Базовый селект
	qb := r.sb.
		Select(tableInboxMessagesColumns...).
		From(tableInboxMessages).
		OrderBy(columnInboxReceivedAt).
		Limit(uint64(o.Limit))

	if o.MaxAttempts > 0 {
		qb = qb.Where(squirrel.Lt{columnInboxAttempts: o.MaxAttempts})
	}

	if len(o.Status) > 0 {
		qb = qb.Where(squirrel.Eq{columnInboxStatus: o.Status})
	}

	// Блокировка строк для конкурентных воркеров
	if o.WithLock {
		qb = qb.Suffix("FOR UPDATE SKIP LOCKED")
	}

	// Выполнение
	conn := r.db.GetQueryEngine(ctx)

	var rows []Inbox
	if err := conn.Selectx(ctx, &rows, qb); err != nil {
		log.Println(err.Error())
		return nil
	}

	// Маппинг в доменную модель
	messages := make([]*model.Inbox, 0, len(rows))
	for i := range rows {
		var processedAt *time.Time
		if rows[i].ProcessedAt.Valid {
			t := rows[i].ProcessedAt.V
			processedAt = &t
		}
		messages = append(messages, &model.Inbox{
			ID:          rows[i].ID,
			Topic:       rows[i].Topic,
			Partition:   rows[i].Partition,
			Offset:      rows[i].Offset,
			Payload:     rows[i].Payload,
			Status:      rows[i].Status,
			Attempts:    rows[i].Attempts,
			LastError:   rows[i].LastError,
			ReceivedAt:  rows[i].ReceivedAt,
			ProcessedAt: processedAt,
		})
	}

	log.Println(messages)

	return messages
}
