package v1

import (
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/models/users"
	"github.com/PechatnovVladimir/msa_big_tech/users/internal/app/usecases/users/dto"
	userPB "github.com/PechatnovVladimir/msa_big_tech/users/pkg/proto/api/users/v1"
)

func dtoSearchByNicknameFromSearchByNicknameRequest(request *userPB.SearchByNicknameRequest) dto.SearchByNickname {
	q := dto.Query{}

	if request.Query.IDs != nil {
		q.IDs = request.Query.IDs
	}

	if request.Query.Email != nil {
		q.Email = request.Query.Email
	}

	if request.Query.Nickname != nil {
		q.Nickname = request.Query.Nickname
	}

	if request.Query.Createdfrom != nil {
		*q.CreatedFrom = (request.Query.Createdfrom).AsTime()
	}

	if request.Query.Createdto != nil {
		*q.CreatedTo = (request.Query.Createdto).AsTime()
	}

	return dto.SearchByNickname{
		Query: q,
		Limit: request.Limit,
	}

}

func responseSearchByNicknameFromUserProfilesModel(p []*users.UserProfile) []*userPB.UserProfile {
	userProfiles := make([]*userPB.UserProfile, len(p))

	for i, _ := range p {
		userProfiles[i] = &userPB.UserProfile{
			UserId:    p[i].ID,
			Nickname:  p[i].Nickname,
			Email:     p[i].Email,
			AvatarUrl: p[i].Avatar,
			Bio:       p[i].Bio,
		}
	}
	return userProfiles
}
