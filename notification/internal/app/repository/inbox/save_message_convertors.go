package inbox

import (
	"database/sql"
	"github.com/PechatnovVladimir/msa_big_tech/notification/internal/app/model"
	"time"
)

func fromModel(in *model.Inbox) *Inbox {
	inbox := &Inbox{
		ID:         in.ID,
		Topic:      in.Topic,
		Partition:  in.Partition,
		Offset:     in.Offset,
		Payload:    in.Payload,
		Status:     in.Status,
		Attempts:   in.Attempts,
		LastError:  in.LastError,
		ReceivedAt: in.ReceivedAt,
		ProcessedAt: func(t *time.Time) sql.Null[time.Time] {
			if in.ProcessedAt != nil {
				return sql.Null[time.Time]{V: *in.ProcessedAt, Valid: true}
			}
			return sql.Null[time.Time]{}
		}(in.ProcessedAt),
	}
	return inbox
}
