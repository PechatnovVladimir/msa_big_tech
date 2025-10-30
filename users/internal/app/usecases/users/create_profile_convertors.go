package users

import (
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
	"time"
)

func modelUserProfileFromCreateProfileDto(d dto.CreateProfile) *users.UserProfile {
	//создаем профайл
	userProfile := users.NewUserProfile()

	userProfile.ID = d.ID
	userProfile.Email = d.Email
	userProfile.Nickname = d.Nickname

	if d.Bio != nil {
		userProfile.Bio = *d.Bio
	}

	if d.Avatar != nil {
		userProfile.Avatar = *d.Avatar
	}

	userProfile.CreateAt = time.Now()

	return userProfile
}
