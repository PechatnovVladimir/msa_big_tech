package main

import (
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/server"
	users "github.com/PechatnovVladimir/msa_big_tech/users/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	authService := server.NewUserService()
	grpcServer := grpc.NewServer()
	users.RegisterUserServiceServer(grpcServer, authService)

	reflection.Register(grpcServer) // регистрируем дополнительные обработчики

	log.Printf("User server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
