package gateway

import (
	"context"
	authpb "github.com/PechatnovVladimir/msa_big_tech/auth/pkg/api"
	chatpb "github.com/PechatnovVladimir/msa_big_tech/chat/pkg/api"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
)

func Rest() {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err := authpb.RegisterAuthServiceHandlerFromEndpoint(ctx, mux, "auth-service:50051", opts)
	if err != nil {
		panic(err)
	}
	err = chatpb.RegisterChatServiceHandlerFromEndpoint(ctx, mux, "chat-service:50052", opts)
	if err != nil {
		panic(err)
	}
	log.Printf("server listening at 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}

}
