package app

import (
	"context"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/adapters/authservice"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/adapters/userservice"
	socialGPRS "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/controllers/social/grpc"
	v1 "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/controllers/social/grpc/v1"
	socialRepo "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/repositories/social"
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context) (err error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	authServiceAdapter := authservice.New()
	userServiceAdapter := userservice.New()
	socialRepository := socialRepo.New()

	socialUseCase := social.New(social.Deps{
		AuthService: authServiceAdapter,
		UserService: userServiceAdapter,
		SocialRepo:  socialRepository,
	})

	//grpc
	grpcServer, err := socialGPRS.New(v1.Deps{
		SocialUseCase: socialUseCase,
	})
	if err != nil {
		return fmt.Errorf("authGRPC.New: %w", err)
	}

	log.Println("Social service started!")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig // wait signal

	grpcServer.Close()

	log.Println("Social service stopped!!!!")

	return nil

}
