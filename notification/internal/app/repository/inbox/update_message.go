package inbox

import (
	"context"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/PechatnovVladimir/msa_big_tech/notification/internal/app/model"
	"github.com/PechatnovVladimir/msa_big_tech/notification/internal/app/usecase/inbox"
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
		Set(columnInboxStatus, o.Status)

	if o.Status == model.StatusProcessing {
		qb = qb.Set(columnInboxAttempts, squirrel.Expr(columnInboxAttempts+" + ?", o.IncAttempts))
	}

	if o.Status == model.StatusProcessed {
		qb = qb.Set(columnInboxProcessedAt, o.ProcessedAt)
	}

	if o.Status == model.StatusFailed {
		qb = qb.Set(columnInboxLastError, o.LastError)
	}

	if len(o.IDs) > 0 {
		qb = qb.Where(squirrel.Eq{columnInboxID: o.IDs}) // id IN (...)
	}

	conn := r.db.GetQueryEngine(ctx)
	if _, err := conn.Execx(ctx, qb); err != nil {
		return fmt.Errorf("%s: %w", api, err)
	}
	return nil
}
