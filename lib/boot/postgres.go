package boot

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/lib/postgres"
	"github.com/PechatnovVladimir/msa_big_tech/lib/postgres/transaction_manager"
	"time"
)

func (app *App) Postgres(ctx context.Context) (*postgres.Connection, *transaction_manager.TransactionManager, error) {
	if app.db == nil {
		conn, err := postgres.NewConnectionPool(ctx, app.cfg.Postgres.DSN(),
			//TODO надо вынести в конфиг наверное нижеследующие параметры подключения
			postgres.WithMaxConnIdleTime(time.Minute),
			postgres.WithMinConnectionsCount(3),
			postgres.WithMaxConnectionsCount(10),
		)

		if err != nil {
			return nil, nil, err
		}
		tx := transaction_manager.New(conn)
		app.db = conn
		app.tx = tx
	}
	return app.db, app.tx, nil
}
