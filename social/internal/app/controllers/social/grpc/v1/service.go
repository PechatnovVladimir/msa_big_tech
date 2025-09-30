package v1

import (
	"github.com/PechatnovVladimir/msa_big_tech/social/internal/app/usecases/social"
	socialPB "github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
)

// зависимости
type Deps struct {
	SocialUseCase social.UseCase
}

type Service struct {
	socialPB.UnimplementedSocialServiceServer
	Deps
}

func New(d Deps) *Service {
	return &Service{
		Deps: d,
	}
}
