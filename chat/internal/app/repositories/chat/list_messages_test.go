package chat

import (
	"context"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/lib/logger"
	"github.com/PechatnovVladimir/msa_big_tech/lib/pagination"
	connection "github.com/PechatnovVladimir/msa_big_tech/lib/postgres"
	"log"
	"testing"
	"time"
)

func TestRepository_ListMessages(t *testing.T) {
	ctx := context.Background()

	//соединение
	conn, err := connection.NewConnectionPool(ctx, DSN(),
		connection.WithMaxConnIdleTime(time.Minute),
		connection.WithMinConnectionsCount(3),
		connection.WithMaxConnectionsCount(10),
	)
	if err != nil {
		logger.Fatal(ctx, err)
	}
	defer conn.Close()

	//менеджер транзакций
	txManager := connection.NewTxManager(conn)

	r := NewRepository(txManager)

	now := time.Time{}

	p := pagination.NewOptions(
		pagination.WithLimit(2),
		pagination.WithCursor(pagination.Cursor{
			ID:   nil,
			Time: &now,
		}))

	messages, err := r.ListMessages(ctx, "d0b49af3-5484-42f6-a316-72cdd68c9d89", p)

	if err != nil {
		t.Fatal(err)
	}

	for _, m := range messages {
		fmt.Println(m)
	}
}
