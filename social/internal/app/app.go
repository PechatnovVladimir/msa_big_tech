package app

import (
	"context"
	"fmt"
	socialGPRS "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/controllers/social/grpc"
	v1 "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/controllers/social/grpc/v1"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context) (err error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	//grpc
	grpcServer, err := socialGPRS.New(v1.Deps{})
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
