package server

import (
	"context"
	"errors"
	auth "github.com/PechatnovVladimir/msa_big_tech/auth/pkg/api"
)

type AuthService struct {
	auth.UnimplementedAuthServiceServer
}

func NewAuthService() *AuthService {
	return &AuthService{}
}

func (s *AuthService) Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	if request.Email == "" || request.Password == "" {
		return &auth.RegisterResponse{}, errors.New("email or password is empty")
	}

	return &auth.RegisterResponse{UserId: 999}, nil
}

func (s *AuthService) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	if request.Email == "" || request.Password == "" {
		return &auth.LoginResponse{}, errors.New("email or password is empty")
	}
	return &auth.LoginResponse{AccessToken: "AccessToken", RefreshToken: "RefreshToken", UserId: 999}, nil
}

func (s *AuthService) Refresh(ctx context.Context, request *auth.RefreshRequest) (*auth.RefreshResponse, error) {
	if request.RefreshToken == "" {
		return &auth.RefreshResponse{}, errors.New("RefreshToken is empty")
	}

	return &auth.RefreshResponse{AccessToken: "AccessToken", RefreshToken: "RefreshToken", UserId: 999}, nil
}
