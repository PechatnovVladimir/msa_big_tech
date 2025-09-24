package main

import (
	"github.com/PechatnovVladimir/msa_big_tech/chat/internal/server"
	chat "github.com/PechatnovVladimir/msa_big_tech/chat/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	chatService := server.NewChatService()
	grpcServer := grpc.NewServer()
	chat.RegisterChatServiceServer(grpcServer, chatService)

	reflection.Register(grpcServer) // регистрируем дополнительные обработчики

	log.Printf("chat server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
