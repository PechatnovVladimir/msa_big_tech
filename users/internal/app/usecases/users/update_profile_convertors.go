package users

import (
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
)

func modelUserProfileForUpdateFromUpdateProfileDto(d dto.UpdateProfile) *users.UserProfileForUpdate {
	userProfile := &users.UserProfileForUpdate{
		ID: d.ID,
	}

	if d.Email != nil {
		userProfile.Email = d.Email
	}

	if d.Nickname != nil {
		userProfile.Nickname = d.Nickname
	}

	if d.Bio != nil {
		userProfile.Bio = d.Bio
	}

	if d.Avatar != nil {
		userProfile.Avatar = d.Avatar
	}

	return userProfile

}
