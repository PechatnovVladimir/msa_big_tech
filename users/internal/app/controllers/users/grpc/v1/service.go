package v1

import (
	usersUC "github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users"
	usersPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
)

type Service struct {
	usersPB.UnimplementedUserServiceServer
	uc *usersUC.Service
}

func New(us *usersUC.Service) *Service {
	return &Service{
		uc: us,
	}
}
