package app

import (
	"context"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/adapters/userservice"
	authGPRS "github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/controllers/auth/grpc"
	v1 "github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/controllers/auth/grpc/v1"
	//authRepo "github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/repositories/auth/inmemory"
	authRepo "github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/repositories/auth"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/usecases/auth"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context) (err error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	userServiceAdapter := userservice.New()

	authRepository := authRepo.New()

	authUseCase := auth.New(auth.Deps{
		UserService: userServiceAdapter,
		AuthRepo:    authRepository,
	})

	//grpc
	grpcServer, err := authGPRS.New(v1.Deps{
		AuthUseCase: authUseCase,
	})
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
