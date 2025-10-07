package app

import (
	"context"
	"fmt"
	usersGPRS "github.com/PechatnovVladimir/msa_big_tech/users/internal/app/controllers/users/grpc"
	userRepo "github.com/PechatnovVladimir/msa_big_tech/users/internal/app/repositories/users"
	usersUC "github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/pkg/postgres/users"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(ctx context.Context) (err error) {

	conn, err := users.NewConnectionPool(ctx, DSN(),
		users.WithMaxConnIdleTime(time.Minute),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	txMngr := users.NewTxManager(conn)

	//репозиторий
	repo := userRepo.NewRepository(txMngr)

	//юзкейс
	uc := usersUC.New(repo)

	//grpc
	grpcServer, err := usersGPRS.New(uc)
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
