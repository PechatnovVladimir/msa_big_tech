package v1

import (
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
	userPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
)

func dtoGetProfileByIDFromGetProfileByIDRequest(request *userPB.GetProfileByIDRequest) dto.GetProfileById {
	result := dto.GetProfileById{
		ID: request.UserId,
	}
	return result
}
