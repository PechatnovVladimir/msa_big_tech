package server

import (
	"context"
	auth "github.com/PechatnovVladimir/msa_big_tech/auth/pkg/proto/api/auth/v1"
	"log"
)

type AuthService struct {
	auth.UnimplementedAuthServiceServer
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	log.Println("AuthService Register called")
	return &auth.RegisterResponse{UserId: 999}, nil
}

func (s *AuthService) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	log.Println("AuthService Login called")
	return &auth.LoginResponse{AccessToken: "AccessToken999", RefreshToken: "RefreshToken999", UserId: 999}, nil
}

func (s *AuthService) Refresh(ctx context.Context, request *auth.RefreshRequest) (*auth.RefreshResponse, error) {
	log.Println("AuthService Refresh called")
	return &auth.RefreshResponse{AccessToken: "AccessToken999", RefreshToken: "RefreshToken999", UserId: 999}, nil
}
