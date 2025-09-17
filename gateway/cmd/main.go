package main

import (
	"context"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"

	authpb "github.com/PechatnovVladimir/msa_big_tech/auth/pkg/api"
)

func main() {
	ctx := context.Background()

	r := chi.NewRouter()

	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	}))

	// Create gRPC gateway multiplexer
	gwMux := runtime.NewServeMux()

	// gRPC connections to microservices
	authConn, err := grpc.Dial("auth-service:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Failed to connect to auth service:", err)
	}
	defer authConn.Close()

	// Register gRPC handlers
	err = authpb.RegisterAuthServiceHandlerClient(ctx, gwMux, authpb.NewAuthServiceClient(authConn))
	if err != nil {
		log.Fatal("Failed to register auth gateway:", err)
	}

}
