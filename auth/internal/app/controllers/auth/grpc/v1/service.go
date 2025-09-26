package v1

import (
	authPB "github.com/PechatnovVladimir/msa_big_tech/auth/pkg/proto/api/auth/v1"
)

type Service struct {
	authPB.UnimplementedAuthServiceServer
	//uc *authUC.Service
}

func New() *Service {
	return &Service{
		//uc: us,
	}
}
