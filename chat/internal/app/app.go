package app

import (
	"context"
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/adapters/userprovider"
	chatGPRS "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/controllers/chat/grpc"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/controllers/chat/grpc/v1"
	chatRepo "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/repositories/chat"
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/usecases/chat"
	connection "github.com/PechatnovVladimir/msa_big_tech/pkg/postgres"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(ctx context.Context) (err error) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	ctx = context.WithValue(ctx, "CurrentUser", os.Getenv("CurrentUser"))

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

	userServiceProvider := userprovider.New()
	chatRepository := chatRepo.NewRepository(txManager)

	chatUseCase := chat.New(chat.Deps{
		UserProvider: userServiceProvider,
		ChatRepo:     chatRepository,
	})

	//grpc
	grpcServer, err := chatGPRS.New(v1.Deps{
		ChatUseCase: chatUseCase,
	})
	if err != nil {
		return fmt.Errorf("chatGRPC.New: %w", err)
	}

	log.Println("Chat service started!")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig // wait signal

	grpcServer.Close()

	log.Println("Chat service stopped!!!!")

	return nil

}
