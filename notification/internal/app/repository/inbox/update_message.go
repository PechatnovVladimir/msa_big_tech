package inbox

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/PechatnovVladimir/msa_big_tech/notification/internal/app/usecase/inbox"
	"log"
	"time"
)

func (r *Repository) UpdateMessages(ctx context.Context, opts ...inbox.UpdateMessageOption) error {
	const api = "outbox.Repository.UpdateEvents"

	o := inbox.CollectUpdateMessageOptions(opts...)

	// защита от noop
	//if o.SetPublishedAt == nil && o.IncRetryBy == 0 && o.SetNextAttemptAt == nil {
	//	return nil
	//}

	qb := r.sb.
		Update(tableInboxMessages).
		Set(columnInboxStatus, "processed").
		Set(columnInboxProcessedAt, time.Now().UTC())
	//Where(squirrel.Eq{columnInboxStatus: "received"}).
	//Where(squirrel.Eq{columnInboxProcessedAt: nil})

	if len(o.IDs) > 0 {
		qb = qb.Where(squirrel.Eq{columnInboxID: o.IDs}) // id IN (...)
	}

	log.Println(qb.ToSql())

	conn := r.db.GetQueryEngine(ctx)
	if _, err := conn.Execx(ctx, qb); err != nil {
		return fmt.Errorf("%s: %w", api, err)
	}
	return nil
}
