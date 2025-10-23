package app

import (
	"context"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/lib/config"
	connection "github.com/PechatnovVladimir/msa_big_tech/lib/postgres"
	"github.com/PechatnovVladimir/msa_big_tech/lib/postgres/transaction_manager"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/adapters/userprovider"
	usersGPRS "github.com/PechatnovVladimir/msa_big_tech/users/internal/app/controllers/users/grpc"
	userRepo "github.com/PechatnovVladimir/msa_big_tech/users/internal/app/repositories/users"
	usersUC "github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func RunNew(ctx context.Context, cfg *config.Config) (err error) {
	if cfg == nil {
		log.Fatal("config is nil")
	}

	//соединение
	conn, err := connection.NewConnectionPool(ctx, cfg.Postgres.DSN(),
		connection.WithMaxConnIdleTime(time.Minute),
		connection.WithMinConnectionsCount(3),
		connection.WithMaxConnectionsCount(10),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	//менеджер транзакций
	txManager := transaction_manager.New(conn)

	userProvider := userprovider.New()

	//репозиторий
	userRepository := userRepo.NewRepository(txManager)

	userUseCase := usersUC.New(usersUC.Deps{
		UserRepo:     userRepository,
		UserProvider: userProvider,
	})

	//grpc
	grpcServer, err := usersGPRS.New(userUseCase, cfg.Grpc)
	if err != nil {
		return fmt.Errorf("%s - grpc.New: %w", cfg.App.Name, err)
	}

	log.Println(fmt.Sprintf("%s started on port %d", cfg.App.Name, cfg.Grpc.Port))

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig // wait signal

	grpcServer.Close()

	log.Println(fmt.Sprintf("%s stopped!", cfg.App.Name))

	return nil

