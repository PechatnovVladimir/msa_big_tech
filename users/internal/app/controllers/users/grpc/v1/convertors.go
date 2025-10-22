package v1

import (
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	userPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
)

func responseUserProfileFromUserProfileModel(profile *users.UserProfile) *userPB.UserProfile {
	result := &userPB.UserProfile{
		UserId:    profile.ID,
		Nickname:  profile.Nickname,
		Email:     profile.Email,
		Bio:       profile.Bio,
		AvatarUrl: profile.Avatar,
	}
	return result
}
