package main

import (
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/server"
	auth "github.com/PechatnovVladimir/msa_big_tech/auth/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	authService := server.NewAuthService()
	grpcServer := grpc.NewServer()
	auth.RegisterAuthServiceServer(grpcServer, authService)

	reflection.Register(grpcServer) // регистрируем дополнительные обработчики

	log.Printf("auth server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
