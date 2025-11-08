package grpc

import (
	"fmt"
	chatGPRS "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/controllers/chat/grpc/v1"
	chatPB "github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/validate"
	"github.com/PechatnovVladimir/msa_big_tech/lib/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"strconv"
)

const (
	defaultGrpcPort = "50052"
)

type Server struct {
	server *grpc.Server
}

func New(d chatGPRS.Deps) (*Server, error) {
	if d.Cfg == nil {
		return nil, fmt.Errorf("missing configuration")
	}

	if d.ChatUseCase == nil {
		return nil, fmt.Errorf("ChatUseCase is not initialized")
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.ChainUnaryInterceptor(validate.Interseptor),
		grpc.ChainUnaryInterceptor(ErrorsUnaryServerInterceptor()),
	)

	chatService := chatGPRS.New(d)
	chatPB.RegisterChatServiceServer(grpcServer, chatService)

	reflection.Register(grpcServer)

	var grpcPort string
	if d.Cfg.Port == 0 {
		grpcPort = defaultGrpcPort
	} else {
		grpcPort = strconv.Itoa(d.Cfg.Port)
	}

	err := start(grpcServer, grpcPort)
	if err != nil {
		return nil, fmt.Errorf("start: %w", err)
	}

	return &Server{server: grpcServer}, nil

}

func start(server *grpc.Server, port string) error {
	conn, err := net.Listen("tcp", net.JoinHostPort("", port))
	if err != nil {
		return fmt.Errorf("net.Listen: %w", err)
	}

	go func() {
		if err := server.Serve(conn); err != nil {
			logger.Info(ctx, "grpc server: Serve")
		}
	}()

	return nil
}

func (s *Server) Close() {
	s.server.GracefulStop()
}
