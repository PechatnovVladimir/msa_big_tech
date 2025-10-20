package inbox

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/notification/internal/app/model"
)

func (r *Repository) SaveMessage(ctx context.Context, msg *model.Inbox) error {

	row := fromModel(msg)

	query := r.sb.Insert(tableInboxMessages).
		Columns(tableInboxMessagesColumns...).
		Values(row.Values(tableInboxMessagesColumns...)...)

	//log.Println(query.ToSql())

	pool := r.db.GetQueryEngine(ctx)

	_, err := pool.Execx(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
