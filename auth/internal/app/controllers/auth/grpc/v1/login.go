package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/auth/pkg/proto/api/auth/v1"
	"log"
)

func (s *Service) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	log.Println("Auth Service Login called")
	return &auth.LoginResponse{AccessToken: "AccessToken999", RefreshToken: "RefreshToken999", UserId: "999"}, nil
}
