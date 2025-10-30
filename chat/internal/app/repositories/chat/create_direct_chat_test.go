package chat

import (
	"context"
	"fmt"
	chat2 "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/models/chat"
	connection "github.com/PechatnovVladimir/msa_big_tech/lib/postgres"
	"github.com/google/uuid"
	"log"
	"testing"
	"time"
)

var (
	PostgresHost     = "localhost"
	PostgresPort     = "5432"
	PostgresDB       = "postgres-chat"
	PostgresUser     = "postgres-chat-user"
	PostgresPassword = "postgres-chat-psw"
)

func DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		PostgresUser, PostgresPassword, PostgresHost, PostgresPort, PostgresDB,
	)
}

func TestRepository_CreateDirectChat(t *testing.T) {
	ctx := context.Background()

	//соединение
	conn, err := connection.NewConnectionPool(ctx, DSN(),
		connection.WithMaxConnIdleTime(time.Minute),
		connection.WithMinConnectionsCount(3),
		connection.WithMaxConnectionsCount(10),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//менеджер транзакций
	txManager := connection.NewTxManager(conn)

	r := NewRepository(txManager)

	chat := chat2.Chat{
		ChatID: uuid.New().String(),
		UserID: uuid.New().String(),
	}

	chatOUT, err := r.CreateDirectChat(ctx, &chat)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(chatOUT)
}
