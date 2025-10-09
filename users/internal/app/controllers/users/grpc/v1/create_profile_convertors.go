package v1

import (
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
	userPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
)

func dtoCreateProfileFromCreateProfileRequest(request *userPB.CreateProfileRequest) dto.CreateProfileDTO {
	result := dto.CreateProfileDTO{
		ID:       request.UserId,
		Nickname: request.Nickname,
		Email:    request.Email,
	}

	if request.Bio != nil {
		result.Bio = request.Bio
	}

	if request.AvatarUrl != nil {
		result.Avatar = request.AvatarUrl
	}
	return result
}
