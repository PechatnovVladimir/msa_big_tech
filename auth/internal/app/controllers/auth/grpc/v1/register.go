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

func (s *Service) Register(ctx context.Context, request *auth.RegisterRequest) (*auth.RegisterResponse, error) {
	log.Println("Auth Service Register called")

	//валидация по proto описанию
	err := protovalidate.Validate(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	inDTO := dto.RegisterInDTO{
		Email:    request.Email,
		Password: request.Password,
	}

	outDTO, err := s.AuthUseCase.Register(ctx, inDTO)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	out := &auth.RegisterResponse{
		UserId: outDTO.UserID,
	}

	log.Println("Auth Service Register success userid=", outDTO.UserID)

	return out, nil
}
