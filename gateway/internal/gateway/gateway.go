package gateway

import (
	"context"
	authpb "github.com/PechatnovVladimir/msa_big_tech/auth/pkg/proto/api/auth/v1"
	chatpb "github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
	"github.com/PechatnovVladimir/msa_big_tech/gateway/swagger"
	socialpb "github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
	userspb "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
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
	muxGRPC := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	//err := authpb.RegisterAuthServiceHandlerFromEndpoint(ctx, muxGRPC, "localhost:50051", opts)
	err := authpb.RegisterAuthServiceHandlerFromEndpoint(ctx, muxGRPC, "auth-service:50051", opts)
	if err != nil {
		panic(err)
	}
	//err = chatpb.RegisterChatServiceHandlerFromEndpoint(ctx, muxGRPC, "localhost:50052", opts)
	err = chatpb.RegisterChatServiceHandlerFromEndpoint(ctx, muxGRPC, "chat-service:50052", opts)
	if err != nil {
		panic(err)
	}
	//err = socialpb.RegisterSocialServiceHandlerFromEndpoint(ctx, muxGRPC, "localhost:50053", opts)
	err = socialpb.RegisterSocialServiceHandlerFromEndpoint(ctx, muxGRPC, "social-service:50053", opts)
	if err != nil {
		panic(err)
	}
	err = userspb.RegisterUserServiceHandlerFromEndpoint(ctx, muxGRPC, "localhost:50054", opts)
	//err = userspb.RegisterUserServiceHandlerFromEndpoint(ctx, muxGRPC, "users-service:50054", opts)
	if err != nil {
		panic(err)
	}

	mux := http.NewServeMux()
	swagger.SetupSwagger(mux)

	mux.Handle("/v1/", muxGRPC)

	log.Printf("server listening at 8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		panic(err)
	}

}
