package grpc

import (
	"fmt"
	chatGPRS "github.com/PechatnovVladimir/msa_big_tech/chat/internal/app/controllers/chat/grpc/v1"
	chatPB "github.com/PechatnovVladimir/msa_big_tech/chat/pkg/proto/api/chat/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	grpcPort = "50052"
)

type Server struct {
	server *grpc.Server
}

func New(d chatGPRS.Deps) (*Server, error) {
	grpcServer := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)

	chatService := chatGPRS.New(d)
	chatPB.RegisterChatServiceServer(grpcServer, chatService)

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

	log.Println("grpc server: AuthService started on port: " + port)

	return nil
}

func (s *Server) Close() {
	s.server.GracefulStop()

	log.Println("grpc server: UserService closed")
}
