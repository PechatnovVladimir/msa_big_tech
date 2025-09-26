package app

import (
	"context"
	"fmt"
	authGPRS "github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/controllers/auth/grpc"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context) (err error) {

	//репозиторий
	//repo := authRepo.New()

	//юзкейс
	//uc := authUC.New(repo)

	//grpc
	grpcServer, err := authGPRS.New()
	if err != nil {
		return fmt.Errorf("authGRPC.New: %w", err)
	}

	log.Println("Auth service started!")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig // wait signal

	grpcServer.Close()

	log.Println("Auth service stopped!!!!")

	return nil

}
