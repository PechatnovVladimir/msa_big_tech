package v1

import (
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
	userPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
)

func dtoUpdateProfileFromUpdateProfileRequest(request *userPB.UpdateProfileRequest) dto.UpdateProfileDTO {
	result := dto.UpdateProfileDTO{
		ID: request.UserId,
	}

	if request.Email != nil {
		result.Email = request.Email
	}

	if request.Nickname != nil {
		result.Nickname = request.Nickname
	}

	if request.Bio != nil {
		result.Bio = request.Bio
	}

	if request.AvatarUrl != nil {
		result.Avatar = request.AvatarUrl
	}

	return result
}
