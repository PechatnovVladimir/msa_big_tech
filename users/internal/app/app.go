package app

import (
	"context"
	"fmt"
	connection "github.com/PechatnovVladimir/msa_big_tech/pkg/postgres"
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

func Run(ctx context.Context) (err error) {

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

	userProvider := userprovider.New()

	//репозиторий
	userRepository := userRepo.NewRepository(txManager)

	userUseCase := usersUC.New(usersUC.Deps{
		UserRepo:     userRepository,
		UserProvider: userProvider,
	})

	//grpc
	grpcServer, err := usersGPRS.New(userUseCase)
	if err != nil {
		return fmt.Errorf("grpc.New: %w", err)
	}

	log.Println("User service started!")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig // wait signal

	grpcServer.Close()

	log.Println("User service stopped!!!!")

	return nil

}
