package grpc

import (
	"fmt"
	"github.com/PechatnovVladimir/msa_big_tech/chat/pkg/validate"
	socialGPRS "github.com/PechatnovVladimir/msa_big_tech/social/internal/app/controllers/social/grpc/v1"
	socialPB "github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	grpcPort = "50053"
)

type Server struct {
	server *grpc.Server
}

func New(d socialGPRS.Deps) (*Server, error) {
	grpcServer := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.ChainUnaryInterceptor(validate.Interseptor),
		grpc.ChainUnaryInterceptor(ErrorsUnaryServerInterceptor()),
	)

	socialService := socialGPRS.New(d)
	socialPB.RegisterSocialServiceServer(grpcServer, socialService)

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
			log.Println("grpc server: Serve")
		}
	}()

	log.Println("grpc server: Social Service started on port: " + port)

	return nil
}

func (s *Server) Close() {
	s.server.GracefulStop()

	log.Println("grpc server: Social Service closed")
}
