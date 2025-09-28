package v1

import (
	"github.com/PechatnovVladimir/msa_big_tech/auth/internal/app/usecases/auth"
	authPB "github.com/PechatnovVladimir/msa_big_tech/auth/pkg/proto/api/auth/v1"
)

type Deps struct {
	AuthUseCase auth.UseCase
}

type Service struct {
	authPB.UnimplementedAuthServiceServer
	Deps
}

func New(d Deps) *Service {
	return &Service{
		Deps: d,
	}
}
