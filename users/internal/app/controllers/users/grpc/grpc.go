package grpc

import (
	"fmt"
	usersGPRS "github.com/PechatnovVladimir/msa_big_tech/users/internal/app/controllers/users/grpc/v1"
	usersUC "github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users"
	usersPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
	"github.com/PechatnovVladimir/msa_big_tech/users/pkg/validate"

	"buf.build/go/protovalidate"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

const (
	grpcPort = "50054"
)

type Server struct {
	server    *grpc.Server
	validator *protovalidate.Validator
}

func New(uc *usersUC.Service) (*Server, error) {
	grpcServer := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
		grpc.ChainUnaryInterceptor(validate.Interseptor),
		grpc.ChainUnaryInterceptor(ErrorsUnaryServerInterceptor()),
	)

	userService := usersGPRS.New(uc)
	usersPB.RegisterUserServiceServer(grpcServer, userService)

	reflection.Register(grpcServer)

	err := start(grpcServer, grpcPort)
	if err != nil {
		return nil, fmt.Errorf("start: %w", err)
	}

	validator, err := protovalidate.New(
		protovalidate.WithDisableLazy(),
		protovalidate.WithMessages(
			&usersPB.CreateProfileRequest{},
			&usersPB.UpdateProfileRequest{},
		),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize validator: %w", err)
	}

	return &Server{
		server:    grpcServer,
		validator: &validator,
	}, nil

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

	log.Println("grpc server: UserService started on port: " + port)

	return nil
}

func (s *Server) Close() {
	s.server.GracefulStop()

	log.Println("grpc server: UserService closed")
}
