package v1

import (
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/auth/pkg/proto/api/auth/v1"
	"log"
)

func (s *Service) Refresh(ctx context.Context, request *auth.RefreshRequest) (*auth.RefreshResponse, error) {
	log.Println("Auth Service Refresh called")
	return &auth.RefreshResponse{AccessToken: "AccessToken999", RefreshToken: "RefreshToken999", UserId: 999}, nil
}
