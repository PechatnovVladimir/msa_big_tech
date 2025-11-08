package grpc

import (
	"context"
	"fmt"
	authGPRS "github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/controllers/auth/grpc/v1"
	authPB "github.com/PechatnovVladimir/msa_big_tech/auth/pkg/proto/api/auth/v1"
	"github.com/PechatnovVladimir/msa_big_tech/lib/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"net"
)

const (
	grpcPort = "50051"
)

type Server struct {
	server *grpc.Server
}

func New(d authGPRS.Deps) (*Server, error) {
	grpcServer := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)

	authService := authGPRS.New(d)
	authPB.RegisterAuthServiceServer(grpcServer, authService)

	reflection.Register(grpcServer)

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
			logger.Info(context.TODO(), "grpc server: Serve")
		}
	}()

	logger.Infof(context.TODO(), "grpc server: AuthService started on port:", port)

	return nil
}

func (s *Server) Close() {
	s.server.GracefulStop()

	logger.Info(context.TODO(), "grpc server: UserService closed")
}
