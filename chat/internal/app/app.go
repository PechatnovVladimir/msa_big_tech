package app

import (
	"context"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/adapters/userservice"
	chatGPRS "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/controllers/chat/grpc"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/controllers/chat/grpc/v1"
	chatRepo "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/repositories/chat"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func Run(ctx context.Context) (err error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	userServiceAdapter := userservice.New()
	chatRepository := chatRepo.New()

	chatUseCase := chat.New(chat.Deps{
		UserService: userServiceAdapter,
		ChatRepo:    chatRepository,
	})

	//grpc
	grpcServer, err := chatGPRS.New(v1.Deps{
		ChatUseCase: chatUseCase,
	})
	if err != nil {
		return fmt.Errorf("authGRPC.New: %w", err)
	}

	log.Println("Chat service started!")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig // wait signal

	grpcServer.Close()

	log.Println("Chat service stopped!!!!")

	return nil

}
