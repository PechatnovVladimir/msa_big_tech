package app

import (
	"context"
	"fmt"
	usersGPRS "github.com/PechatnovVladimir/msa_big_tech/users/internal/app/controllers/users/grpc"
	userRepo "github.com/PechatnovVladimir/msa_big_tech/users/internal/app/repositories/inmemory/users"
	usersUC "github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context) (err error) {

	//репозиторий
	repo := userRepo.New(5)

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
