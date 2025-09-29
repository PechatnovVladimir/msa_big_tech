package v1

import (
	"buf.build/go/protovalidate"
	"context"
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/usecases/auth/dto"
	"github.com/PechatnovVladimir/msa_big_tech/auth/pkg/proto/api/auth/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Service) Login(ctx context.Context, request *auth.LoginRequest) (*auth.LoginResponse, error) {
	log.Println("Auth Service Login called")
	//валидация по proto описанию
	err := protovalidate.Validate(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	itDTO := dto.LoginInDTO{
		Email:    request.Email,
		Password: request.Password,
	}

	outDTO, err := s.AuthUseCase.Login(ctx, itDTO)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	out := &auth.LoginResponse{
		AccessToken:  outDTO.AccessToken,
		RefreshToken: outDTO.RefreshToken,
		UserId:       outDTO.UserID,
	}

	return out, nil
}
