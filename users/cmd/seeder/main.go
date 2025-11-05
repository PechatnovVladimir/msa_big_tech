package main

import (
	"context"
	connection "github.com/PechatnovVladimir/msa_big_tech/lib/postgres"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	userRepo "github.com/PechatnovVladimir/msa_big_tech/users/internal/app/repositories/users"
	"github.com/brianvoe/gofakeit/v7"
	"github.com/google/uuid"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	//соединение
	conn, err := connection.NewConnectionPool(ctx, app.DSN(),
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

	//репозиторий
	repo := userRepo.NewRepository(txManager)

	for range 100 {
		p := users.UserProfile{
			ID:       uuid.New().String(),
			Email:    gofakeit.Email(),
			Nickname: gofakeit.FirstName(),
			Bio:      gofakeit.Name(),
			Avatar:   gofakeit.URL(),
			CreateAt: time.Time{},
		}
		_, err := repo.CreateProfile(ctx, &p)
		if err != nil {
			log.Println(err)
		}
	}

}
