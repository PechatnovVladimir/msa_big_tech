package v1

import (
	socialPB "github.com/PechatnovVladimir/msa_big_tech/social/pkg/proto/api/social/v1"
)

// зависимости
type Deps struct {
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
