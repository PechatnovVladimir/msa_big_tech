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

func (s *Service) Refresh(ctx context.Context, request *auth.RefreshRequest) (*auth.RefreshResponse, error) {
	log.Println("Auth Service Refresh called")
	//валидация по proto описанию
	err := protovalidate.Validate(request)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	inDTO := dto.RefreshInDTO{
		RefreshToken: request.RefreshToken,
	}

	outDTO, err := s.AuthUseCase.Refresh(ctx, inDTO)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	out := &auth.RefreshResponse{
		RefreshToken: outDTO.RefreshToken,
		AccessToken:  outDTO.AccessToken,
		UserId:       outDTO.UserID,
	}
	return out, nil
}
