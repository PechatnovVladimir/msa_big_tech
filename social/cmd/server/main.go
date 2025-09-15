package main

import (
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/server"
	social "github.com/PechatnovVladimir/msa_big_tech/social/pkg/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	socialService := server.NewSocialService()
	grpcServer := grpc.NewServer()
	social.RegisterSocialServiceServer(grpcServer, socialService)

	reflection.Register(grpcServer) // регистрируем дополнительные обработчики

	log.Printf("Social server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
