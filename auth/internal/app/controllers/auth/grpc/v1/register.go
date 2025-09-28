package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/auth/pkg/proto/api/auth/v1"
	"log"
)

func (s *Service) Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	log.Println("Auth Service Register called")
	return &auth.RegisterResponse{UserId: "999"}, nil
}
