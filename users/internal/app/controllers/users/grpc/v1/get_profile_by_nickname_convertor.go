package v1

import (
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
	userPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
)

func dtoGetProfileByNicknameFromGetProfileByNicknameRequest(request *userPB.GetProfileByNicknameRequest) dto.GetProfileByNickname {
	result := dto.GetProfileByNickname{
		Nickname: request.Nickname,
	}
	return result
}
